import { useCallback, useState } from "react"

export function usePermissions() {
  const [isRequesting, setIsRequesting] = useState(false)

  const requestAudioPermission = useCallback(async (): Promise<boolean> => {
    if (isRequesting) return false

    setIsRequesting(true)
    try {
      // Check if getUserMedia is supported
      if (!navigator.mediaDevices || !navigator.mediaDevices.getUserMedia) {
        console.error("getUserMedia not supported")
        return false
      }

      // First check if permission is already granted
      try {
        const result = await navigator.permissions.query({ name: "microphone" as PermissionName })
        if (result.state === "granted") {
          return true
        }
        if (result.state === "denied") {
          alert("麦克风权限被拒绝，请在地浏览器设置中开启麦克风权限")
          return false
        }
      } catch {
        // permissions API not supported, try directly
      }

      // Request permission by trying to get the stream
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true })
      // Stop the stream immediately - we just wanted permission
      stream.getTracks().forEach((track) => track.stop())
      return true
    } catch (err) {
      console.error("Failed to get audio permission:", err)
      alert("无法获取麦克风权限，请在地浏览器设置中开启麦克风权限")
      return false
    } finally {
      setIsRequesting(false)
    }
  }, [isRequesting])

  return { requestAudioPermission }
}
