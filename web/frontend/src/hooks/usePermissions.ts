import { useCallback, useState } from 'react';

export function usePermissions() {
  const [permissionGranted, setPermissionGranted] = useState<boolean | null>(null);

  const requestAudioPermission = useCallback(async (): Promise<boolean> => {
    try {
      const result = await navigator.permissions.query({ name: 'microphone' as PermissionName });
      if (result.state === 'granted') {
        setPermissionGranted(true);
        return true;
      }
      
      if (result.state === 'prompt') {
        const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
        stream.getTracks().forEach(track => track.stop());
        setPermissionGranted(true);
        return true;
      }
      
      setPermissionGranted(false);
      return false;
    } catch {
      // 浏览器不支持 permissions API，尝试直接获取
      try {
        const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
        stream.getTracks().forEach(track => track.stop());
        setPermissionGranted(true);
        return true;
      } catch {
        setPermissionGranted(false);
        return false;
      }
    }
  }, []);

  return { requestAudioPermission, permissionGranted };
}
