import { useCallback, useRef, useState } from "react"

interface UseRecorderReturn {
  isRecording: boolean
  startRecording: () => Promise<void>
  stopRecording: () => Promise<string | null>
  clearAudio: () => void
}

export function useRecorder(): UseRecorderReturn {
  const [isRecording, setIsRecording] = useState(false)
  const mediaRecorderRef = useRef<MediaRecorder | null>(null)
  const chunksRef = useRef<Blob[]>([])
  const streamRef = useRef<MediaStream | null>(null)

  const startRecording = useCallback(async () => {
    if (isRecording) return
    
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true })
      streamRef.current = stream
      
      const mediaRecorder = new MediaRecorder(stream, {
        mimeType: "audio/webm;codecs=opus",
      })
      mediaRecorderRef.current = mediaRecorder
      chunksRef.current = []

      mediaRecorder.ondataavailable = (event) => {
        if (event.data.size > 0) {
          chunksRef.current.push(event.data)
        }
      }

      mediaRecorder.start()
      setIsRecording(true)
    } catch (error) {
      console.error("Failed to start recording:", error)
      throw error
    }
  }, [isRecording])

  const stopRecording = useCallback(async (): Promise<string | null> => {
    return new Promise((resolve) => {
      const mediaRecorder = mediaRecorderRef.current
      
      if (!mediaRecorder || mediaRecorder.state === "inactive") {
        setIsRecording(false)
        resolve(null)
        return
      }

      mediaRecorder.onstop = () => {
        const blob = new Blob(chunksRef.current, { type: "audio/webm" })
        const reader = new FileReader()
        
        reader.onloadend = () => {
          const base64 = reader.result as string
          setIsRecording(false)
          
          // Stop all tracks
          if (streamRef.current) {
            streamRef.current.getTracks().forEach((track) => track.stop())
            streamRef.current = null
          }
          
          resolve(base64)
        }
        
        reader.onerror = () => {
          setIsRecording(false)
          resolve(null)
        }
        
        reader.readAsDataURL(blob)
      }

      mediaRecorder.stop()
    })
  }, [])

  const clearAudio = useCallback(() => {
    chunksRef.current = []
    if (streamRef.current) {
      streamRef.current.getTracks().forEach((track) => track.stop())
      streamRef.current = null
    }
    setIsRecording(false)
  }, [])

  return {
    isRecording,
    startRecording,
    stopRecording,
    clearAudio,
  }
}
