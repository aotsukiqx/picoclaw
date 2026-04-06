// components/CaptureButton.tsx
import React from 'react';

interface CaptureButtonProps {
  onPhoto: () => void;
  onVideo: () => void;
  disabled?: boolean;
}

export function CaptureButton({ onPhoto, onVideo, disabled = false }: CaptureButtonProps) {
  return (
    <div className="flex gap-3">
      <button
        onClick={onPhoto}
        disabled={disabled}
        className="flex-1 py-3 px-6 bg-purple-600 hover:bg-purple-700 disabled:bg-gray-400 text-white rounded-lg font-medium transition-colors"
        title="拍照"
      >
        📷 拍照
      </button>
      <button
        onClick={onVideo}
        disabled={disabled}
        className="flex-1 py-3 px-6 bg-red-600 hover:bg-red-700 disabled:bg-gray-400 text-white rounded-lg font-medium transition-colors"
        title="录制3秒视频"
      >
        🎬 3秒视频
      </button>
    </div>
  );
}
