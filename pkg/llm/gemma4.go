package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Gemma4Provider struct {
	apiBase   string
	apiKey    string
	model     string
	maxTokens int
	httpClient *http.Client
}

func NewGemma4Provider(apiBase, apiKey, model string, maxTokens int) *Gemma4Provider {
	if maxTokens <= 0 {
		maxTokens = 8192
	}
	return &Gemma4Provider{
		apiBase:   apiBase,
		apiKey:    apiKey,
		model:     model,
		maxTokens: maxTokens,
		httpClient: &http.Client{Timeout: 120 * time.Second},
	}
}

func (p *Gemma4Provider) Name() string { return "gemma4" }

func (p *Gemma4Provider) Complete(ctx context.Context, prompt string) (string, error) {
	url := p.apiBase + "/v1/chat/completions"
	
	messages := []map[string]string{
		{"role": "user", "content": prompt},
	}
	
	body := map[string]interface{}{
		"model": p.model,
		"messages": messages,
		"max_tokens": p.maxTokens,
		"temperature": 0.7,
	}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Gemma4 API 请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Gemma4 API 错误 (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("Gemma4 响应解析失败: %w", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("Gemma4 返回空响应")
	}
	
	return result.Choices[0].Message.Content, nil
}

func (p *Gemma4Provider) StreamComplete(ctx context.Context, prompt string, onChunk func(string)) error {
	url := p.apiBase + "/v1/chat/completions"
	
	messages := []map[string]string{
		{"role": "user", "content": prompt},
	}
	
	body := map[string]interface{}{
		"model": p.model,
		"messages": messages,
		"max_tokens": p.maxTokens,
		"temperature": 0.7,
		"stream": true,
	}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Gemma4 API 请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Gemma4 API 错误 (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	decoder := json.NewDecoder(resp.Body)
	for {
		var chunk struct {
			Choices []struct {
				Delta struct {
					Content string `json:"content"`
				} `json:"delta"`
			} `json:"choices"`
		}
		if err := decoder.Decode(&chunk); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("Gemma4 流式解析失败: %w", err)
		}
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			onChunk(chunk.Choices[0].Delta.Content)
		}
	}
	
	return nil
}
