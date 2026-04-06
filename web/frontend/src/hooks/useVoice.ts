import { useState, useCallback, useRef, useEffect } from "react"

export interface UseVoiceOptions {
  asrServerUrl: string
  onTranscriptionResult?: (text: string) => void
  onError?: (error: Error) => void
}

export function useVoice(options: UseVoiceOptions) {
  const [isRecording, setIsRecording] = useState(false)
  const [recordingPath, setRecordingPath] = useState<string | null>(null)
  const [isConnected, setIsConnected] = useState(false)
  const mediaRecorderRef = useRef<MediaRecorder | null>(null)
  const audioChunksRef = useRef<Blob[]>([])
  const wsRef = useRef<WebSocket | null>(null)
  const reconnectTimeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null)

  // WebSocket connection with auto-reconnect
  const connectWebSocket = useCallback(() => {
    if (wsRef.current?.readyState === WebSocket.OPEN) return

    try {
      const ws = new WebSocket(options.asrServerUrl)

      ws.onopen = () => {
        setIsConnected(true)
      }

      ws.onclose = () => {
        setIsConnected(false)
        // Auto-reconnect after 3 seconds
        reconnectTimeoutRef.current = setTimeout(() => {
          connectWebSocket()
        }, 3000)
      }

      ws.onerror = () => {
        options.onError?.(new Error("WebSocket connection error"))
      }

      ws.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data)
          if (data.text && options.onTranscriptionResult) {
            options.onTranscriptionResult(data.text)
          }
        } catch {
          // Ignore non-JSON messages
        }
      }

      wsRef.current = ws
    } catch (err) {
      options.onError?.(err instanceof Error ? err : new Error("Failed to connect"))
    }
  }, [options.asrServerUrl, options.onError, options.onTranscriptionResult])

  useEffect(() => {
    connectWebSocket()
    return () => {
      if (reconnectTimeoutRef.current) {
        clearTimeout(reconnectTimeoutRef.current)
      }
      if (wsRef.current) {
        wsRef.current.close()
      }
    }
  }, [connectWebSocket])

  const startRecord = useCallback(async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true })
      const mediaRecorder = new MediaRecorder(stream, {
        mimeType: "audio/webm;codecs=opus",
      })

      audioChunksRef.current = []

      mediaRecorder.ondataavailable = (event) => {
        if (event.data.size > 0) {
          audioChunksRef.current.push(event.data)
        }
      }

      mediaRecorder.onstop = async () => {
        const audioBlob = new Blob(audioChunksRef.current, { type: "audio/webm" })
        const path = URL.createObjectURL(audioBlob)
        setRecordingPath(path)

        // Send to ASR server via WebSocket
        if (wsRef.current?.readyState === WebSocket.OPEN) {
          const reader = new FileReader()
          reader.onloadend = () => {
            const base64 = reader.result as string
            wsRef.current?.send(
              JSON.stringify({
                action: "audio",
                data: base64.split(",")[1], // Remove data URL prefix
                format: "webm",
              })
            )
          }
          reader.readAsDataURL(audioBlob)
        }

        // Stop all tracks
        stream.getTracks().forEach((track) => track.stop())
      }

      mediaRecorderRef.current = mediaRecorder
      mediaRecorder.start(1000) // Collect data every second
      setIsRecording(true)
    } catch (err) {
      options.onError?.(err instanceof Error ? err : new Error("Failed to start recording"))
      throw err
    }
  }, [options.onError])

  const stopRecord = useCallback(async () => {
    return new Promise<string>((resolve, reject) => {
      if (!mediaRecorderRef.current || mediaRecorderRef.current.state === "inactive") {
        reject(new Error("Not recording"))
        return
      }

      const mediaRecorder = mediaRecorderRef.current

      mediaRecorder.onstop = () => {
        const audioBlob = new Blob(audioChunksRef.current, { type: "audio/webm" })
        const path = URL.createObjectURL(audioBlob)
        setRecordingPath(path)
        setIsRecording(false)

        // Send finish signal to WebSocket
        if (wsRef.current?.readyState === WebSocket.OPEN) {
          wsRef.current.send(JSON.stringify({ action: "finish" }))
        }

        resolve(path)
      }

      mediaRecorder.stop()
    })
  }, [])

  const sendJsonMessage = useCallback((message: object) => {
    if (wsRef.current?.readyState === WebSocket.OPEN) {
      wsRef.current.send(JSON.stringify(message))
    }
  }, [])

  return {
    isRecording,
    isConnected,
    recordingPath,
    startRecord,
    stopRecord,
    sendJsonMessage,
  }
}
