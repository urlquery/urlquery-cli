package api

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	DefaultUrlqueryAPI string        = "https://api.urlquery.net"
	DefaultHTTPTimeout time.Duration = 30 * time.Second
)

// Default Argus Client
// NewClient should not return any error with called without any options, and therefore safe to ignore
var DefaultClient, _ = NewClient()

type Client interface {
	NewRequest(method string, path string, body io.Reader) (*http.Request, error)
	Do(req *http.Request) (*http.Response, error)
	DoRequest(method string, path string, body io.Reader) (*http.Response, error)
}

type OptionsClientFunc func(client *httpClient) error

// httpClient represents the REST API client.
type httpClient struct {
	baseURL string
	client  *http.Client

	headerProperty map[string]string
	apiKey         string
}

func NewClient(opts ...OptionsClientFunc) (*httpClient, error) {
	client := &httpClient{
		baseURL: DefaultUrlqueryAPI,
		client: &http.Client{
			Timeout: DefaultHTTPTimeout,
		},
		headerProperty: make(map[string]string),
	}

	for _, opt := range opts {
		err := opt(client)
		if err != nil {
			return nil, err
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

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Add custom headers to request
	for k, v := range c.headerProperty {
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
