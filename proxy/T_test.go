package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReverseProxy(t *testing.T) {
	// Create a new router and reverse proxy
	r := chi.NewRouter()
	proxy := NewReverseProxy("hugo", "1313")
	r.Use(proxy.ReverseProxy)

	// Create a test server
	server := httptest.NewServer(r)
	defer server.Close()

	// Create a request to the test server
	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body for the root path
	expectedRootResponse := "Hello from API"
	if body := w.Body.String(); body != expectedRootResponse {
		t.Errorf("Unexpected response body for root path: got %v want %v", body, expectedRootResponse)
	}

	// Create a request for the "/api/" path
	reqAPI, err := http.NewRequest("GET", server.URL+"/api/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder for the API request
	wAPI := httptest.NewRecorder()

	// Perform the API request
	r.ServeHTTP(wAPI, reqAPI)

	// Check the status code for the API request
	if status := wAPI.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code for API request: got %v want %v", status, http.StatusOK)
	}

	// Check the response body for the "/api/" path
	expectedAPIResponse := "Hello from API"
	if body := wAPI.Body.String(); body != expectedAPIResponse {
		t.Errorf("Unexpected response body for API path: got %v want %v", body, expectedAPIResponse)
	}
}
