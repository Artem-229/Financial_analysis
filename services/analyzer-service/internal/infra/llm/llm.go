package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"financial_assistant/services/analyzer-service/internal/config"
	"financial_assistant/services/analyzer-service/internal/entities"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type LLM struct {
	config     *config.Config
	httpClient *http.Client
}

func NewLLM(conf *config.Config) *LLM {
	return &LLM{
		config:     conf,
		httpClient: &http.Client{Timeout: conf.Ollama.Timeout},
	}
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

type chatResponse struct {
	Message chatMessage `json:"message"`
}

// Analyze sends each user's per-class transaction summary to Ollama, paired
// with the configured system prompt, and returns one Outcome per user.
func (m *LLM) Analyze(ctx context.Context, transactions map[uuid.UUID]map[string]entities.SortedTransactions) ([]entities.Outcome, error) {
	outcomes := make([]entities.Outcome, 0, len(transactions))

	for userID, byClass := range transactions {
		analysis, err := m.analyzeUser(ctx, byClass)
		if err != nil {
			return nil, fmt.Errorf("analyzing user %s: %w", userID, err)
		}

		outcomes = append(outcomes, entities.Outcome{
			ID:       uuid.New(),
			UserID:   userID,
			Analysis: analysis,
		})
	}

	return outcomes, nil
}

func (m *LLM) analyzeUser(ctx context.Context, byClass map[string]entities.SortedTransactions) (string, error) {
	payload, err := json.Marshal(byClass)
	if err != nil {
		return "", fmt.Errorf("marshalling transactions: %w", err)
	}

	reqBody, err := json.Marshal(chatRequest{
		Model:  m.config.Ollama.Model,
		Stream: false,
		Messages: []chatMessage{
			{Role: "system", Content: m.config.SystemPrompt},
			{Role: "user", Content: string(payload)},
		},
	})
	if err != nil {
		return "", fmt.Errorf("marshalling request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, m.config.Ollama.Host+"/api/chat", bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("calling ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ollama returned status %d", resp.StatusCode)
	}

	var chatResp chatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", fmt.Errorf("decoding ollama response: %w", err)
	}

	return chatResp.Message.Content, nil
}
