package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Client_6281dc is an HTTP client wrapper
type Client_6281dc struct {
	BaseURL    string
	HTTPClient *http.Client
	ProjectID  string
}

// Response_6281dc holds the API response
type Response_6281dc struct {
	Status    int                    `json:"status"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Timestamp string                 `json:"timestamp"`
	ProjectID string                 `json:"project_id"`
}

// NewClient_6281dc creates a new HTTP client
func NewClient_6281dc(baseURL string) *Client_6281dc {
	return &Client_6281dc{
		BaseURL:   baseURL,
		ProjectID: "6281dc",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Get_6281dc performs a GET request
func (c *Client_6281dc) Get_6281dc(path string) (*Response_6281dc, error) {
	url := c.BaseURL + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("User-Agent", "go-client/1.0")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	_ = json.Unmarshal(body, &data)

	return &Response_6281dc{
		Status:    resp.StatusCode,
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
		ProjectID: c.ProjectID,
	}, nil
}

func main() {
	base := os.Getenv("API_BASE")
	if base == "" {
		base = "https://httpbin.org"
	}
	client := NewClient_6281dc(base)
	result, err := client.Get_6281dc("/get")
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
