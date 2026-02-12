package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Pricing struct {
	Prompt            string          `json: "prompt"`
	Completion        string          `json: "completion"`
	Request           json.RawMessage `json: "request"`
	Image             json.RawMessage `json: "image"`
	ImageToken        json.RawMessage `json: "image_token"`
	ImageOutput       json.RawMessage `json: "image_output"`
	Audio             json.RawMessage `json: "audio"`
	AudioOutput       json.RawMessage `json: "audio_output"`
	InputAudioCache   json.RawMessage `json: "input_audio_cache"`
	WebSearch         json.RawMessage `json: "web_search"`
	InternalReasoning json.RawMessage `json: "intenal_reasoning"`
	InputCacheRead    json.RawMessage `json: ""`
	InputCacheWrite   json.RawMessage `json: ""`
	Discount          json.RawMessage `json: ""`
}

type Model struct {
	Id                  string
	CanonicalSlug       string
	Name                string
	Created             float64
	Pricing             any
	ContextLength       float64
	Architecture        any
	TopProvider         any
	PerRequestLimits    any
	SupportedParameters any
	DefaultParameters   string
	Description         string
	ExpirationDate      string
}

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

func (c *OpenRouterClient) Models() error {
	url := "https://openrouter.ai/api/v1/models"
	res, err := c.makeRequest("GET", url)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))

	return nil
}

func (c *OpenRouterClient) makeRequest(method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
