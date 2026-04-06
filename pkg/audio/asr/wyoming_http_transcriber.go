package asr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// WyomingHTTPTranscriber 实现基于 HTTP 的异步语音识别
type WyomingHTTPTranscriber struct {
	serverURL  string
	apiKey     string
	model      string
	language   string
	httpClient *http.Client
}

// NewWyomingHTTPTranscriber 创建新的 Wyoming HTTP 转写器
func NewWyomingHTTPTranscriber(serverURL, apiKey, model, language string) *WyomingHTTPTranscriber {
	return &WyomingHTTPTranscriber{
		serverURL: serverURL,
		apiKey:    apiKey,
		model:     model,
		language:  language,
		httpClient: &http.Client{Timeout: 120 * time.Second},
	}
}

// Transcribe 执行 HTTP 异步语音识别
func (t *WyomingHTTPTranscriber) Transcribe(ctx context.Context, audioFilePath string) (*TranscriptionResponse, error) {
	// 读取音频文件
	audioData, err := io.ReadFile(audioFilePath)
	if err != nil {
		return nil, fmt.Errorf("读取音频文件失败：%w", err)
	}

	// 创建上传请求
	url := t.serverURL + "/v1/audio/transcriptions"
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(audioData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败：%w", err)
	}

	req.Header.Set("Content-Type", "audio/wav")
	if t.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+t.apiKey)
	}

	// 发送请求
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP 请求失败：%w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ASR 服务错误 (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	// 解析响应
	var result struct {
		Text     string  `json:"text"`
		Language string  `json:"language"`
		Duration float64 `json:"duration"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &TranscriptionResponse{
		Text:     result.Text,
		Language: result.Language,
		Duration: result.Duration,
	}, nil
}
