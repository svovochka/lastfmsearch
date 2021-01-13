package lastfm

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client LasFM API client
type Client struct {
	Endpoint   string
	ApiKey     string
	httpClient *http.Client
	timeout    int
}

// NewClient Create new client
func NewClient(endpoint string, apiKey string, timeout int) *Client {
	return &Client{
		Endpoint:   endpoint,
		ApiKey:     apiKey,
		httpClient: &http.Client{},
		timeout:    timeout,
	}
}

// Query Returns API response
func (c *Client) Query(ctx context.Context, q string, s interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.timeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, q, nil)
	if err != nil {
		return fmt.Errorf("failed to create http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute http request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code `%d`", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}
