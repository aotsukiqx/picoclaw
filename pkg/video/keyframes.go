package video

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// KeyFrameExtractor 关键帧提取器
type KeyFrameExtractor struct {
	ffmpegPath string
	outputDir  string
}

// NewKeyFrameExtractor 创建新的关键帧提取器
func NewKeyFrameExtractor(outputDir string) *KeyFrameExtractor {
	return &KeyFrameExtractor{
		ffmpegPath: "ffmpeg",
		outputDir:  outputDir,
	}
}

// ExtractKeyFrames 从视频中提取关键帧（均匀采样，每秒1帧）
func (e *KeyFrameExtractor) ExtractKeyFrames(ctx context.Context, videoPath string) ([]string, error) {
	if _, err := os.Stat(videoPath); err != nil {
		return nil, fmt.Errorf("视频文件不存在: %w", err)
	}

	if err := os.MkdirAll(e.outputDir, 0755); err != nil {
		return nil, fmt.Errorf("创建输出目录失败: %w", err)
	}

	filename := filepath.Base(videoPath)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	outputPattern := filepath.Join(e.outputDir, name+"_frame_%03d.jpg")

	// 使用 ffmpeg 提取帧：每秒1帧，最多5帧
	args := []string{
		"-i", videoPath,
		"-vf", "fps=1",
		"-q:v", "2",
		"-frames:v", "5",
		outputPattern,
	}

	cmd := exec.CommandContext(ctx, e.ffmpegPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("ffmpeg 提取帧失败: %w (output: %s)", err, string(output))
	}

	// 查找生成的帧文件
	pattern := filepath.Join(e.outputDir, name+"_frame_*.jpg")
	frames, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("查找帧文件失败: %w", err)
	}

	return frames, nil
}

// ExtractKeyFrameAtTime 在指定时间提取单帧
func (e *KeyFrameExtractor) ExtractKeyFrameAtTime(ctx context.Context, videoPath string, timestamp time.Duration) (string, error) {
	if _, err := os.Stat(videoPath); err != nil {
		return "", fmt.Errorf("视频文件不存在: %w", err)
	}

	if err := os.MkdirAll(e.outputDir, 0755); err != nil {
		return "", fmt.Errorf("创建输出目录失败: %w", err)
	}

	filename := filepath.Base(videoPath)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	outputPath := filepath.Join(e.outputDir, fmt.Sprintf("%s_frame_at_%s.jpg", name, timestamp.String()))

	args := []string{
		"-ss", timestamp.String(),
		"-i", videoPath,
		"-frames:v", "1",
		"-q:v", "2",
		outputPath,
	}

	cmd := exec.CommandContext(ctx, e.ffmpegPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ffmpeg 提取单帧失败: %w (output: %s)", err, string(output))
	}

	return outputPath, nil
}

// ExtractFramesBySceneChange 基于场景变化检测提取关键帧
func (e *KeyFrameExtractor) ExtractFramesBySceneChange(ctx context.Context, videoPath string, threshold float64) ([]string, error) {
	if _, err := os.Stat(videoPath); err != nil {
		return nil, fmt.Errorf("视频文件不存在: %w", err)
	}

	if err := os.MkdirAll(e.outputDir, 0755); err != nil {
		return nil, fmt.Errorf("创建输出目录失败: %w", err)
	}

	filename := filepath.Base(videoPath)
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	outputPattern := filepath.Join(e.outputDir, name+"_scene_%03d.jpg")

	if threshold <= 0 {
		threshold = 0.3
	}

	// 使用 select 滤镜检测场景变化
	args := []string{
		"-i", videoPath,
		"-vf", fmt.Sprintf("select='gt(scene,%f)',showinfo", threshold),
		"-vsync", "vfr",
		"-q:v", "2",
		outputPattern,
	}

	cmd := exec.CommandContext(ctx, e.ffmpegPath, args...)
	_, err := cmd.CombinedOutput()
	// 场景变化检测可能返回非零状态，忽略

	pattern := filepath.Join(e.outputDir, name+"_scene_*.jpg")
	frames, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("查找帧文件失败: %w", err)
	}

	return frames, nil
}
