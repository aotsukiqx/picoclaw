import React, { useState, useRef } from 'react';

interface VoiceButtonProps {
  onPressIn: () => void;
  onPressOut: () => void;
  disabled?: boolean;
  isConnected?: boolean;
}

export function VoiceButton({ onPressIn, onPressOut, disabled = false, isConnected = true }: VoiceButtonProps) {
  const [isPressed, setIsPressed] = useState(false);
  const scaleAnim = useRef(1);
  const isDisabled = disabled || !isConnected;

  const handlePressIn = () => {
    if (isDisabled) return;
    setIsPressed(true);
    onPressIn();
  };

  const handlePressOut = () => {
    if (isDisabled) return;
    setIsPressed(false);
    onPressOut();
  };

  const buttonStyle: React.CSSProperties = {
    width: 56,
    height: 56,
    borderRadius: 28,
    backgroundColor: isPressed ? '#6D28D9' : isDisabled ? '#9CA3AF' : '#7C3AED',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    cursor: isDisabled ? 'not-allowed' : 'pointer',
    transform: `scale(${scaleAnim.current})`,
    transition: 'background-color 0.2s, transform 0.1s',
    boxShadow: isPressed ? 'none' : '0 4px 6px rgba(124, 58, 237, 0.3)',
  };

  return (
    <div
      onMouseDown={handlePressIn}
      onMouseUp={handlePressOut}
      onMouseLeave={handlePressOut}
      onTouchStart={handlePressIn}
      onTouchEnd={handlePressOut}
      style={buttonStyle}
      role="button"
      aria-label={isDisabled ? '语音输入（未连接）' : '语音输入'}
      aria-disabled={isDisabled}
    >
      {isConnected ? (
        <svg width="24" height="24" fill="none" stroke="white" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11a7 7 0 01-7 7m0 0a7 7 0 01-7-7m7 7v4m0 0H8m4 0h4m-4-8a3 3 0 01-3-3V5a3 3 0 116 0v6a3 3 0 01-3 3z" />
        </svg>
      ) : (
        <svg width="24" height="24" fill="none" stroke="white" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z" />
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2" />
        </svg>
      )}
    </div>
  );
}
