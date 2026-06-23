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

// Client_97e796 is an HTTP client wrapper
type Client_97e796 struct {
	BaseURL    string
	HTTPClient *http.Client
	ProjectID  string
}

// Response_97e796 holds the API response
type Response_97e796 struct {
	Status    int                    `json:"status"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Timestamp string                 `json:"timestamp"`
	ProjectID string                 `json:"project_id"`
}

// NewClient_97e796 creates a new HTTP client
func NewClient_97e796(baseURL string) *Client_97e796 {
	return &Client_97e796{
		BaseURL:   baseURL,
		ProjectID: "97e796",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Get_97e796 performs a GET request
func (c *Client_97e796) Get_97e796(path string) (*Response_97e796, error) {
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

	return &Response_97e796{
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
	client := NewClient_97e796(base)
	result, err := client.Get_97e796("/get")
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
