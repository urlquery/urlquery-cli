package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		opts    []OptionsClientFunc
		wantErr bool
	}{
		{
			name:    "default client",
			opts:    nil,
			wantErr: false,
		},
		{
			name:    "client with API key",
			opts:    []OptionsClientFunc{ApiKey("test-key")},
			wantErr: false,
		},
		{
			name:    "client with custom base URL",
			opts:    []OptionsClientFunc{ApiGWBase("https://custom.api.com")},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && client == nil {
				t.Error("NewClient() returned nil client")
			}
		})
	}
}

func TestHttpClient_NewRequest(t *testing.T) {
	client, err := NewClient(ApiKey("test-key"))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req, err := client.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("NewRequest() error = %v", err)
	}

	if req.Header.Get("x-apikey") != "test-key" {
		t.Error("API key header not set correctly")
	}

	if req.Header.Get("User-Agent") != DefaultUserAgent {
		t.Error("User-Agent header not set correctly")
	}

	if req.Header.Get("Accept") != "application/json" {
		t.Error("Accept header not set correctly")
	}
}

func TestApiKey(t *testing.T) {
	client := &httpClient{}
	opt := ApiKey("test-key")

	err := opt(client)
	if err != nil {
		t.Errorf("ApiKey() error = %v", err)
	}

	if client.apiKey != "test-key" {
		t.Errorf("ApiKey not set correctly, got %s, want test-key", client.apiKey)
	}
}

func TestApiGWBase(t *testing.T) {
	client := &httpClient{}
	testURL := "https://custom.api.com"
	opt := ApiGWBase(testURL)

	err := opt(client)
	if err != nil {
		t.Errorf("ApiGWBase() error = %v", err)
	}

	if client.baseURL != testURL {
		t.Errorf("BaseURL not set correctly, got %s, want %s", client.baseURL, testURL)
	}
}

func TestHttpClient_DoRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client, err := NewClient(ApiGWBase(server.URL), ApiKey("test-key"))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	resp, err := client.DoRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("DoRequest() error = %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestDefaultClient(t *testing.T) {
	if DefaultClient == nil {
		t.Error("DefaultClient should not be nil")
	}

	// Test that DefaultClient has expected defaults
	req, err := DefaultClient.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("DefaultClient.NewRequest() error = %v", err)
	}

	if req.URL.Host != "api.urlquery.net" {
		t.Errorf("Expected default host api.urlquery.net, got %s", req.URL.Host)
	}
}

func TestClientTimeout(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test that client was created successfully
	if client == nil {
		t.Error("Expected non-nil client")
	}
}
