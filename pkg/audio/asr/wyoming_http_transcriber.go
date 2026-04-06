package asr

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	asrPollInterval = 2 * time.Second
	asrMaxTimeout   = 20 * time.Minute
	asrMaxRetries   = 600
)

// WyomingHTTPTranscriber implements asynchronous ASR via Wyoming HTTP protocol (port 10380).
type WyomingHTTPTranscriber struct {
	apiKey     string
	apiBase    string
	modelID    string
	httpClient *http.Client
}

// NewWyomingHTTPTranscriber creates an HTTP-based Wyoming ASR transcriber.
func NewWyomingHTTPTranscriber(apiKey, apiBase, modelID string) *WyomingHTTPTranscriber {
	return &WyomingHTTPTranscriber{
		apiKey:     apiKey,
		apiBase:    strings.TrimRight(apiBase, "/"),
		modelID:    modelID,
		httpClient: &http.Client{Timeout: 25 * time.Minute},
	}
}

// Name returns the provider name.
func (t *WyomingHTTPTranscriber) Name() string {
	return "wyoming-http"
}

// Transcribe runs the full three-step async workflow: create task, poll for
// completion, then retrieve the result.
func (t *WyomingHTTPTranscriber) Transcribe(ctx context.Context, audioFilePath string) (*TranscriptionResponse, error) {
	jobID, err := t.createTask(ctx, audioFilePath)
	if err != nil {
		return nil, fmt.Errorf("create transcription task failed: %w", err)
	}

	if err := t.waitForCompletion(ctx, jobID); err != nil {
		return nil, fmt.Errorf("transcription processing failed: %w", err)
	}

	return t.getResult(ctx, jobID)
}

// createTask POST /v1/audio/transcriptions
func (t *WyomingHTTPTranscriber) createTask(ctx context.Context, audioFilePath string) (string, error) {
	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to open audio file: %w", err)
	}
	defer audioFile.Close()

	body := &strings.Builder{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(audioFilePath))
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(part, audioFile); err != nil {
		return "", err
	}

	writer.WriteField("model", t.modelID)
	writer.WriteField("response_format", "json")
	writer.Close()

	url := t.apiBase + "/v1/audio/transcriptions"
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(body.String()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("create task failed (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var jobResp struct {
		ID string `json:"id"`
	}
	json.Unmarshal(bodyBytes, &jobResp)
	return jobResp.ID, nil
}

// waitForCompletion GET /v1/audio/transcriptions/{id} — polls until done.
func (t *WyomingHTTPTranscriber) waitForCompletion(ctx context.Context, jobID string) error {
	url := fmt.Sprintf("%s/v1/audio/transcriptions/%s", t.apiBase, jobID)
	deadline := time.Now().Add(asrMaxTimeout)
	retries := 0

	for time.Now().Before(deadline) && retries < asrMaxRetries {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(asrPollInterval):
		}

		req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
		resp, err := t.httpClient.Do(req)
		if err != nil {
			continue
		}
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var statusResp struct {
			Status string `json:"status"`
			Error  string `json:"error"`
		}
		json.Unmarshal(bodyBytes, &statusResp)

		switch statusResp.Status {
		case "completed":
			return nil
		case "failed":
			return fmt.Errorf("transcription failed: %s", statusResp.Error)
		case "cancelled":
			return fmt.Errorf("transcription cancelled")
		}
		retries++
	}
	return fmt.Errorf("transcription timeout (exceeded %v)", asrMaxTimeout)
}

// getResult GET /v1/audio/transcriptions/{id}/result
func (t *WyomingHTTPTranscriber) getResult(ctx context.Context, jobID string) (*TranscriptionResponse, error) {
	url := fmt.Sprintf("%s/v1/audio/transcriptions/%s/result", t.apiBase, jobID)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get result failed (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	var result TranscriptionResponse
	json.Unmarshal(bodyBytes, &result)
	return &result, nil
}
