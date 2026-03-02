package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const defaultAPIURL = "https://api.openai.com/v1/chat/completions"

// Ask gửi câu hỏi lên OpenAI Chat Completions, trả về nội dung trả lời.
func Ask(ctx context.Context, apiKey, userMessage string) (string, error) {
	if apiKey == "" {
		return "", fmt.Errorf("openai: api key chưa cấu hình")
	}
	body := map[string]any{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "system", "content": "Bạn là trợ lý thân thiện, trả lời ngắn gọn bằng tiếng Việt."},
			{"role": "user", "content": userMessage},
		},
		"max_tokens": 1024,
	}
	raw, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, defaultAPIURL, bytes.NewReader(raw))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("openai: status %d", resp.StatusCode)
	}
	var out struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	if len(out.Choices) == 0 {
		return "", fmt.Errorf("openai: không có câu trả lời")
	}
	return out.Choices[0].Message.Content, nil
}
