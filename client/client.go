package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Client FastGPT客户端
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// NewClient 创建新的FastGPT客户端
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// DoRequest 发送HTTP请求
func (c *Client) DoRequest(method, path string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, c.BaseURL+path, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "go-fastgpt-client")

	return c.HTTPClient.Do(req)
}

// ParseResponse 解析HTTP响应
func (c *Client) ParseResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}
