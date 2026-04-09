import { useState, useCallback, useRef, useEffect } from 'react';

export interface UseVoiceOptions {
  asrServerUrl: string;
  onTranscriptionResult?: (text: string) => void;
  onError?: (error: Error) => void;
}

export function useVoice(options: UseVoiceOptions) {
  const [isRecording, setIsRecording] = useState(false);
  const [_recordingPath, _setRecordingPath] = useState<string | null>(null); // recordingPath used for debug
  const [isConnected, setIsConnected] = useState(false);
  const wsRef = useRef<WebSocket | null>(null);
  const mediaRecorderRef = useRef<MediaRecorder | null>(null);
  const audioChunksRef = useRef<Blob[]>([]);

  const connect = useCallback(() => {
    if (wsRef.current?.readyState === WebSocket.OPEN) return;

    const ws = new WebSocket(options.asrServerUrl);
    
    ws.onopen = () => setIsConnected(true);
    ws.onclose = () => setIsConnected(false);
    ws.onerror = () => options.onError?.(new Error('WebSocket 连接失败'));
    
    ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        if (data.text) {
          options.onTranscriptionResult?.(data.text);
        }
      } catch {
        options.onTranscriptionResult?.(event.data);
      }
    };

    wsRef.current = ws;
  }, [options]);

  const disconnect = useCallback(() => {
    wsRef.current?.close();
    wsRef.current = null;
    setIsConnected(false);
  }, []);

  const startRecord = useCallback(async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      const mediaRecorder = new MediaRecorder(stream);
      audioChunksRef.current = [];

      mediaRecorder.ondataavailable = (event) => {
        if (event.data.size > 0) {
          audioChunksRef.current.push(event.data);
        }
      };

      mediaRecorder.onstop = async () => {
        const audioBlob = new Blob(audioChunksRef.current, { type: 'audio/webm' });
        const arrayBuffer = await audioBlob.arrayBuffer();
        wsRef.current?.send(arrayBuffer);
        stream.getTracks().forEach(track => track.stop());
      };

      mediaRecorder.start();
      mediaRecorderRef.current = mediaRecorder;
      setIsRecording(true);
      connect();
    } catch (err) {
      options.onError?.(err as Error);
    }
  }, [connect, options]);

  const stopRecord = useCallback(async () => {
    if (mediaRecorderRef.current && isRecording) {
      mediaRecorderRef.current.stop();
      setIsRecording(false);
    }
  }, [isRecording]);

  useEffect(() => {
    return () => {
      disconnect();
      if (mediaRecorderRef.current) {
        mediaRecorderRef.current.stop();
      }
    };
  }, [disconnect]);

  return {
    isRecording,
    isConnected,
    startRecord,
    stopRecord,
  };
}
