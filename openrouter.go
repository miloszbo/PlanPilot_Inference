package main

import (
	"errors"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://openrouter.ai/api/v1/chat/completions"

type OpenRouterClient struct {
	httpClient *http.Client
	apiKey     string
}

func NewOpenRouterClient() (*OpenRouterClient, error) {
	key := os.Getenv("OPENROUTER_API_KEY")
	if key == "" {
		return nil, errors.New("OPENROUTER_API_KEY not set")
	}

	return &OpenRouterClient{
		apiKey: key,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}, nil
}

func (c *OpenRouterClient) Chat() error {
	return nil
}
