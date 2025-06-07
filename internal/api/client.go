package api

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultUrlqueryAPI string        = "https://api.urlquery.net"
	DefaultHTTPTimeout time.Duration = 30 * time.Second
	DefaultUserAgent   string        = "urlquery-cli/1.0 (+https://github.com/urlquery/urlquery-cli)"
)

// Default Client
// NewClient should not return any error when called without any options, and therefore safe to ignore
var DefaultClient, _ = NewClient()

// Client defines the interface for HTTP operations
type Client interface {
	NewRequest(method string, path string, body io.Reader) (*http.Request, error)
	NewRequestWithContext(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error)
	Do(req *http.Request) (*http.Response, error)
	DoRequest(method string, path string, body io.Reader) (*http.Response, error)
	DoRequestWithContext(ctx context.Context, method string, path string, body io.Reader) (*http.Response, error)
}

type OptionsClientFunc func(client *httpClient) error

// httpClient represents the REST API client.
type httpClient struct {
	baseURL string
	client  *http.Client

	headers   map[string]string
	apiKey    string
	userAgent string
}

func NewClient(opts ...OptionsClientFunc) (*httpClient, error) {
	client := &httpClient{
		baseURL:   DefaultUrlqueryAPI,
		userAgent: DefaultUserAgent,
		client: &http.Client{
			Timeout: DefaultHTTPTimeout,
		},
		headers: make(map[string]string),
	}

	for _, opt := range opts {
		err := opt(client)
		if err != nil {
			return nil, fmt.Errorf("failed to apply client option: %w", err)
		}
	}
	return client, nil
}

// API Key Authentication
func ApiKey(key string) OptionsClientFunc {
	return func(client *httpClient) error {
		client.apiKey = key
		return nil
	}
}

// Custom API Gateway
func ApiGWBase(apigw_base string) OptionsClientFunc {
	return func(client *httpClient) error {
		client.baseURL = apigw_base
		return nil
	}
}

func (c *httpClient) NewRequest(method string, path string, body io.Reader) (*http.Request, error) {
	url := c.baseURL + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if c.apiKey != "" {
		req.Header.Add("x-apikey", c.apiKey)
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Add custom headers to request
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	return req, nil
}

// NewRequestWithContext creates a new HTTP request with context
func (c *httpClient) NewRequestWithContext(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error) {
	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if c.apiKey != "" {
		req.Header.Add("x-apikey", c.apiKey)
	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Add custom headers to request
	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

	return req, nil
}

// Do executes a HTTP request.
func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

// DoRequest makes an HTTP request and executes it.
func (c *httpClient) DoRequest(method string, path string, body io.Reader) (*http.Response, error) {
	req, err := c.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// DoRequestWithContext makes an HTTP request with context and executes it.
func (c *httpClient) DoRequestWithContext(ctx context.Context, method string, path string, body io.Reader) (*http.Response, error) {
	req, err := c.NewRequestWithContext(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// DecodeResponse decodes the HTTP response body.
func DecodeResponse(resp *http.Response, target any) error {

	err := handleResponseError(resp)
	if err != nil {
		return err
	}

	err = decodeResponseBody(resp, target)
	if err != nil {
		target = nil
	}

	return err
}

func decodeResponseBody(resp *http.Response, target any) error {
	var err error

	if resp.Body == nil || target == nil {
		return nil // Nothing to decode
	}
	defer resp.Body.Close()

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		var reader *gzip.Reader
		reader, err = gzip.NewReader(resp.Body)
		if err == nil {
			defer reader.Close()
			err = json.NewDecoder(reader).Decode(target)
		}

	default:
		err = json.NewDecoder(resp.Body).Decode(target)
	}

	return err
}
