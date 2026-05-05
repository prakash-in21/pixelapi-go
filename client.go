package pixelapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

// Client is a PixelAPI HTTP client.
type Client struct {
	APIKey  string
	BaseURL string
	HTTP    *http.Client
}

// NewClient creates a PixelAPI client. Pass your API key from pixelapi.dev/app.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.pixelapi.dev",
		HTTP:    &http.Client{Timeout: 120 * time.Second},
	}
}

// Result is the output of any successful PixelAPI call.
type Result struct {
	OutputURL   string  `json:"output_url"`
	CreditsUsed float64 `json:"credits_used"`
	JobID       string  `json:"job_id,omitempty"`
}

func (c *Client) post(endpoint string, body url.Values) (*Result, error) {
	req, err := http.NewRequest("POST", c.BaseURL+endpoint, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-API-Key", c.APIKey)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == 401 {
		return nil, errors.New("invalid API key")
	}
	if resp.StatusCode == 402 {
		return nil, errors.New("insufficient credits")
	}
	if resp.StatusCode == 429 {
		return nil, errors.New("rate limit exceeded")
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("PixelAPI error %d: %s", resp.StatusCode, string(respBody))
	}
	var r Result
	if err := json.Unmarshal(respBody, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

// Generate creates an AI image from a text prompt.
func (c *Client) Generate(prompt string) (*Result, error) {
	v := url.Values{}
	v.Set("prompt", prompt)
	return c.post("/v1/image/generate", v)
}

// RemoveBackground removes the background from an image URL.
func (c *Client) RemoveBackground(imageURL string) (*Result, error) {
	v := url.Values{}
	v.Set("image_url", imageURL)
	return c.post("/v1/image/remove-background", v)
}

// Upscale 4x upscales an image at imageURL.
func (c *Client) Upscale(imageURL string) (*Result, error) {
	v := url.Values{}
	v.Set("image_url", imageURL)
	v.Set("scale", "4")
	return c.post("/v1/image/upscale", v)
}

// FaceRestore restores faces in a photo at imageURL.
func (c *Client) FaceRestore(imageURL string) (*Result, error) {
	v := url.Values{}
	v.Set("image_url", imageURL)
	return c.post("/v1/image/face-restore", v)
}

// Save downloads a Result's output_url to a local file path.
func (c *Client) Save(r *Result, path string) error {
	if r == nil || r.OutputURL == "" {
		return errors.New("nil or empty result")
	}
	resp, err := http.Get(r.OutputURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}
