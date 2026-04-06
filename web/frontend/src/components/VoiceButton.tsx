import React, { useState, useRef } from "react"

import { Mic, MicOff } from "lucide-react"

import { cn } from "@/lib/utils"

interface VoiceButtonProps {
  onPressIn: () => void
  onPressOut: () => void
  disabled?: boolean
  isConnected?: boolean
  isRecording?: boolean
  className?: string
}

export function VoiceButton({
  onPressIn,
  onPressOut,
  disabled = false,
  isConnected = true,
  isRecording = false,
  className,
}: VoiceButtonProps) {
  const [isPressed, setIsPressed] = useState(false)
  const scaleAnim = useRef<number>(1)

  const isDisabled = disabled || !isConnected

  const handlePressIn = () => {
    if (isDisabled) return
    setIsPressed(true)
    // Simple scale animation via CSS
    scaleAnim.current = 1.2
    onPressIn()
  }

  const handlePressOut = () => {
    setIsPressed(false)
    scaleAnim.current = 1
    onPressOut()
  }

  return (
    <button
      type="button"
      onMouseDown={handlePressIn}
      onMouseUp={handlePressOut}
      onMouseLeave={isPressed ? handlePressOut : undefined}
      onTouchStart={handlePressIn}
      onTouchEnd={handlePressOut}
      disabled={isDisabled}
      aria-label={isDisabled ? "语音输入（未连接）" : isRecording ? "松开结束录音" : "按住说话"}
      aria-pressed={isRecording}
      className={cn(
        "relative inline-flex shrink-0 items-center justify-center rounded-full transition-all duration-150",
        "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
        "disabled:pointer-events-none disabled:opacity-50",
        isPressed && "scale-110",
        className
      )}
      style={{
        width: 56,
        height: 56,
        transform: `scale(${scaleAnim.current})`,
      }}
    >
      {/* Pulse ring animation when recording */}
      {isRecording && (
        <span
          className={cn(
            "absolute inset-0 animate-ping rounded-full opacity-75",
            "bg-primary text-primary-foreground"
          )}
        />
      )}

      {/* Main button */}
      <span
        className={cn(
          "relative z-10 flex h-full w-full items-center justify-center rounded-full",
          "transition-colors duration-150",
          isDisabled
            ? "bg-muted text-muted-foreground"
            : isRecording
              ? "bg-destructive text-destructive-foreground animate-pulse"
              : isPressed
                ? "bg-primary/90 text-primary-foreground"
                : "bg-primary text-primary-foreground hover:bg-primary/80"
        )}
      >
        {isDisabled ? (
          <MicOff className="h-6 w-6" />
        ) : (
          <Mic className={cn("h-6 w-6", isRecording && "animate-pulse")} />
        )}
      </span>
    </button>
  )
}
