package asr

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WyomingWSTranscriber implements streaming ASR via Wyoming WebSocket protocol (port 10388).
type WyomingWSTranscriber struct {
	apiKey     string
	wsEndpoint string // ws://host:10388/
	httpClient *http.Client
}

// NewWyomingWSTranscriber creates a WebSocket-based Wyoming ASR transcriber.
func NewWyomingWSTranscriber(apiKey, host string, port int) *WyomingWSTranscriber {
	endpoint := fmt.Sprintf("ws://%s:%d/", strings.TrimRight(host, ":"), port)
	return &WyomingWSTranscriber{
		apiKey:     apiKey,
		wsEndpoint: endpoint,
		httpClient: &http.Client{Timeout: 60 * time.Second},
	}
}

// Name returns the provider name.
func (t *WyomingWSTranscriber) Name() string {
	return "wyoming-ws"
}

// TranscribeStream streams audio chunks to the Wyoming WebSocket endpoint and
// delivers interim and final results via callbacks.
func (t *WyomingWSTranscriber) TranscribeStream(
	ctx context.Context,
	audioChunkReader io.Reader, // 20ms @ 16kHz = 640 bytes per chunk
	onInterimResult func(string),
	onFinalResult func(string),
	onError func(error),
) error {
	header := http.Header{}
	if t.apiKey != "" {
		header.Set("Authorization", "Bearer "+t.apiKey)
	}

	conn, _, err := websocket.DefaultDialer.Dial(t.wsEndpoint, header)
	if err != nil {
		return fmt.Errorf("WebSocket connection failed: %w", err)
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	// Send audio data asynchronously.
	go func() {
		defer wg.Done()
		buf := make([]byte, 640) // 20ms @ 16kHz * 2 bytes = 640 bytes
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n, err := audioChunkReader.Read(buf)
				if err == io.EOF {
					conn.WriteMessage(websocket.TextMessage, []byte(`{"action":"finish"}`))
					return
				}
				if err != nil {
					onError(fmt.Errorf("audio read failed: %w", err))
					return
				}
				if err := conn.WriteMessage(websocket.BinaryMessage, buf[:n]); err != nil {
					onError(fmt.Errorf("audio send failed: %w", err))
					return
				}
			}
		}
	}()

	// Receive recognition results.
	go func() {
		defer wg.Done()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					onError(fmt.Errorf("WebSocket read failed: %w", err))
				}
				return
			}

			var result struct {
				Type   string `json:"type"`
				Result struct {
					Text string `json:"text"`
				} `json:"result"`
			}

			if err := json.Unmarshal(msg, &result); err != nil {
				continue
			}

			switch result.Type {
			case "connection":
				// Connection acknowledgement, ignore.
			case "result":
				if result.Result.Text != "" {
					onFinalResult(result.Result.Text)
				}
			case "interim":
				if onInterimResult != nil && result.Result.Text != "" {
					onInterimResult(result.Result.Text)
				}
			case "error":
				onError(fmt.Errorf("recognition error: %s", string(msg)))
			}
		}
	}()

	wg.Wait()
	return nil
}
