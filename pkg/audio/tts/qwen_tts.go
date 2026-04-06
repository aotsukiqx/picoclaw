package tts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type QwenTTSProvider struct {
	apiBase    string
	voice      string
	language   string
	speedRatio float64
	httpClient *http.Client
}

func NewQwenTTSProvider(apiBase, voice, language string, speedRatio float64) *QwenTTSProvider {
	if speedRatio <= 0 {
		speedRatio = 1.0
	}
	return &QwenTTSProvider{
		apiBase:    strings.TrimRight(apiBase, "/"),
		voice:      voice,
		language:   language,
		speedRatio: speedRatio,
		httpClient: &http.Client{Timeout: 60 * time.Second},
	}
}

func (p *QwenTTSProvider) Name() string {
	return "qwen-tts"
}

func (p *QwenTTSProvider) Synthesize(ctx context.Context, text string) (io.ReadCloser, error) {
	url := p.apiBase + "/api/v1/tts/synthesize"
	body := map[string]interface{}{
		"text":        text,
		"voice":       p.voice,
		"language":    p.language,
		"speed_ratio": p.speedRatio,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Qwen TTS 服务不可用: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return resp.Body, nil
	case http.StatusTooManyRequests: // 429 限流
		resp.Body.Close()
		return nil, fmt.Errorf("TTS 限流，请稍后重试")
	case http.StatusServiceUnavailable: // 503
		resp.Body.Close()
		return nil, fmt.Errorf("TTS 服务不可用")
	default:
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("TTS 错误 (status %d): %s", resp.StatusCode, string(bodyBytes))
	}
}
