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

// Client_918dcc is an HTTP client wrapper
type Client_918dcc struct {
	BaseURL    string
	HTTPClient *http.Client
	ProjectID  string
}

// Response_918dcc holds the API response
type Response_918dcc struct {
	Status    int                    `json:"status"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Timestamp string                 `json:"timestamp"`
	ProjectID string                 `json:"project_id"`
}

// NewClient_918dcc creates a new HTTP client
func NewClient_918dcc(baseURL string) *Client_918dcc {
	return &Client_918dcc{
		BaseURL:   baseURL,
		ProjectID: "918dcc",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Get_918dcc performs a GET request
func (c *Client_918dcc) Get_918dcc(path string) (*Response_918dcc, error) {
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

	return &Response_918dcc{
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
	client := NewClient_918dcc(base)
	result, err := client.Get_918dcc("/get")
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
