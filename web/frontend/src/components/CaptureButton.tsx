import React from 'react';

interface CaptureButtonProps {
  onPhoto: () => void;
  onVideo: () => void;
  disabled?: boolean;
}

export function CaptureButton({ onPhoto, onVideo, disabled = false }: CaptureButtonProps) {
  const buttonBaseStyle: React.CSSProperties = {
    flex: 1,
    padding: '12px 24px',
    borderRadius: '0.5rem',
    fontWeight: 500,
    transition: 'background-color 0.2s',
    border: 'none',
    cursor: disabled ? 'not-allowed' : 'pointer',
  };

  const photoButtonStyle: React.CSSProperties = {
    ...buttonBaseStyle,
    backgroundColor: disabled ? '#9CA3AF' : '#7C3AED',
    color: 'white',
  };

  const videoButtonStyle: React.CSSProperties = {
    ...buttonBaseStyle,
    backgroundColor: disabled ? '#9CA3AF' : '#DC2626',
    color: 'white',
  };

  const containerStyle: React.CSSProperties = {
    display: 'flex',
    gap: '12px',
  };

  return (
    <div style={containerStyle}>
      <button
        onClick={onPhoto}
        disabled={disabled}
        style={photoButtonStyle}
        title="拍照"
      >
        📷 拍照
      </button>
      <button
        onClick={onVideo}
        disabled={disabled}
        style={videoButtonStyle}
        title="录制3秒视频"
      >
        🎬 3秒视频
      </button>
    </div>
  );
}
