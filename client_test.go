package golang

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const testAPIKey = "7500|V7Yz2LcScGI2HtfbGNc08e2huNI6XOv6ppR7tc7L"

func TestNewClient(t *testing.T) {
	client := NewClient(testAPIKey)

	if client.APIKey != testAPIKey {
		t.Errorf("Expected API key %s, got %s", testAPIKey, client.APIKey)
	}

	if client.BaseURL != DefaultBaseURL {
		t.Errorf("Expected base URL %s, got %s", DefaultBaseURL, client.BaseURL)
	}

	if client.HTTPClient.Timeout != DefaultTimeout {
		t.Errorf("Expected timeout %v, got %v", DefaultTimeout, client.HTTPClient.Timeout)
	}
}

func TestNewClientWithURL(t *testing.T) {
	customURL := "https://custom.api.url"
	client := NewClientWithURL(testAPIKey, customURL)

	if client.APIKey != testAPIKey {
		t.Errorf("Expected API key %s, got %s", testAPIKey, client.APIKey)
	}

	if client.BaseURL != customURL {
		t.Errorf("Expected base URL %s, got %s", customURL, client.BaseURL)
	}
}

func TestSetTimeout(t *testing.T) {
	client := NewClient(testAPIKey)
	newTimeout := 60 * time.Second

	client.SetTimeout(newTimeout)

	if client.HTTPClient.Timeout != newTimeout {
		t.Errorf("Expected timeout %v, got %v", newTimeout, client.HTTPClient.Timeout)
	}
}

func TestDoRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check authorization header
		auth := r.Header.Get("Authorization")
		expectedAuth := "Bearer " + testAPIKey
		if auth != expectedAuth {
			t.Errorf("Expected Authorization header %s, got %s", expectedAuth, auth)
		}

		// Check accept header
		accept := r.Header.Get("Accept")
		if accept != "application/json" {
			t.Errorf("Expected Accept header application/json, got %s", accept)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client := NewClientWithURL(testAPIKey, server.URL)

	resp, err := client.doRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestHandleResponse_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "message": "test"}`))
	}))
	defer server.Close()

	client := NewClientWithURL(testAPIKey, server.URL)

	resp, err := client.doRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	err = client.handleResponse(resp, &result)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !result.Success {
		t.Error("Expected success to be true")
	}

	if result.Message != "test" {
		t.Errorf("Expected message 'test', got '%s'", result.Message)
	}
}

func TestHandleResponse_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))
	}))
	defer server.Close()

	client := NewClientWithURL(testAPIKey, server.URL)

	resp, err := client.doRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	var result struct{}
	err = client.handleResponse(resp, &result)
	if err == nil {
		t.Error("Expected an error for bad request")
	}
}
