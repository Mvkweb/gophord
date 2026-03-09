// Package rest provides a REST client for the Discord API.
//
// This package handles HTTP requests to Discord's REST API, including
// rate limiting, authentication, and error handling.
package rest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Mvkweb/gophord/pkg/json"
	"github.com/Mvkweb/gophord/pkg/types"
)

// Client is a REST client for the Discord API.
type Client struct {
	// Token is the bot token for authentication.
	token string
	// HTTPClient is the underlying HTTP client.
	httpClient *http.Client
	// BaseURL is the base URL for API requests.
	baseURL string
	// UserAgent is the User-Agent header value.
	userAgent string
	// rateLimiter handles rate limiting.
	rateLimiter *RateLimiter
}

// ClientOption is a function that configures a Client.
type ClientOption func(*Client)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithBaseURL sets a custom base URL.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithUserAgent sets a custom User-Agent.
func WithUserAgent(ua string) ClientOption {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// New creates a new REST client with the given bot token.
func New(token string, opts ...ClientOption) *Client {
	c := &Client{
		token:       token,
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		baseURL:     types.BaseURL,
		userAgent:   "DiscordBot (https://github.com/Mvkweb/gophord, 1.0.0)",
		rateLimiter: NewRateLimiter(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Request performs an HTTP request to the Discord API.
func (c *Client) Request(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	return c.RequestWithQuery(ctx, method, path, body, nil)
}

// RequestWithQuery performs an HTTP request with query parameters.
func (c *Client) RequestWithQuery(ctx context.Context, method, path string, body interface{}, query map[string]string) ([]byte, error) {
	// Wait for rate limit
	if err := c.rateLimiter.Wait(ctx, path); err != nil {
		return nil, fmt.Errorf("rate limit wait: %w", err)
	}

	// Build URL
	url := c.baseURL + path
	if len(query) > 0 {
		url += "?"
		first := true
		for k, v := range query {
			if !first {
				url += "&"
			}
			url += k + "=" + v
			first = false
		}
	}

	// Prepare body
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bot "+c.token)
	req.Header.Set("User-Agent", c.userAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	// Update rate limits
	c.rateLimiter.Update(path, resp.Header)

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	// Handle errors
	if resp.StatusCode >= 400 {
		return nil, c.handleError(resp.StatusCode, respBody)
	}

	return respBody, nil
}

// handleError creates an appropriate error from an error response.
func (c *Client) handleError(statusCode int, body []byte) error {
	var apiErr APIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return &HTTPError{
			StatusCode: statusCode,
			Message:    string(body),
		}
	}
	apiErr.StatusCode = statusCode
	return &apiErr
}

// APIError represents a Discord API error.
type APIError struct {
	StatusCode int                    `json:"-"`
	Code       int                    `json:"code"`
	Message    string                 `json:"message"`
	Errors     map[string]interface{} `json:"errors,omitempty"`
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if len(e.Errors) > 0 {
		return fmt.Sprintf("discord api error %d: %s (code: %d, errors: %v)", e.StatusCode, e.Message, e.Code, e.Errors)
	}
	return fmt.Sprintf("discord api error %d: %s (code: %d)", e.StatusCode, e.Message, e.Code)
}

// HTTPError represents a non-API HTTP error.
type HTTPError struct {
	StatusCode int
	Message    string
}

// Error implements the error interface.
func (e *HTTPError) Error() string {
	return fmt.Sprintf("http error %d: %s", e.StatusCode, e.Message)
}

// RateLimiter handles Discord API rate limiting.
type RateLimiter struct {
	mu      sync.Mutex
	buckets map[string]*bucket
	global  *time.Time
}

type bucket struct {
	remaining int
	reset     time.Time
	mu        sync.Mutex
}

// NewRateLimiter creates a new rate limiter.
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		buckets: make(map[string]*bucket),
	}
}

// Wait blocks until the rate limit allows a request.
func (r *RateLimiter) Wait(ctx context.Context, path string) error {
	r.mu.Lock()

	// Check global rate limit
	if r.global != nil && time.Now().Before(*r.global) {
		waitTime := time.Until(*r.global)
		r.mu.Unlock()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(waitTime):
		}

		r.mu.Lock()
	}

	// Get or create bucket
	b, ok := r.buckets[path]
	if !ok {
		b = &bucket{remaining: 1}
		r.buckets[path] = b
	}
	r.mu.Unlock()

	b.mu.Lock()
	defer b.mu.Unlock()

	// Wait if rate limited
	if b.remaining <= 0 && time.Now().Before(b.reset) {
		waitTime := time.Until(b.reset)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(waitTime):
		}
	}

	return nil
}

// Update updates rate limit info from response headers.
func (r *RateLimiter) Update(path string, headers http.Header) {
	// Check for global rate limit
	if headers.Get("X-RateLimit-Global") == "true" {
		if retryAfter := headers.Get("Retry-After"); retryAfter != "" {
			if seconds, err := strconv.ParseFloat(retryAfter, 64); err == nil {
				globalReset := time.Now().Add(time.Duration(seconds * float64(time.Second)))
				r.mu.Lock()
				r.global = &globalReset
				r.mu.Unlock()
			}
		}
		return
	}

	// Update bucket
	r.mu.Lock()
	b, ok := r.buckets[path]
	if !ok {
		b = &bucket{}
		r.buckets[path] = b
	}
	r.mu.Unlock()

	b.mu.Lock()
	defer b.mu.Unlock()

	if remaining := headers.Get("X-RateLimit-Remaining"); remaining != "" {
		if r, err := strconv.Atoi(remaining); err == nil {
			b.remaining = r
		}
	}

	if reset := headers.Get("X-RateLimit-Reset"); reset != "" {
		if ts, err := strconv.ParseFloat(reset, 64); err == nil {
			b.reset = time.Unix(int64(ts), int64((ts-float64(int64(ts)))*1e9))
		}
	}
}
