package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReverseProxy(t *testing.T) {
	r := chi.NewRouter()
	proxy := NewReverseProxy("hugo", "1313")
	r.Use(proxy.ReverseProxy)

	server := httptest.NewServer(r)
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", status, http.StatusOK)
	}

	expectedRootResponse := "Hello from API"
	if body := w.Body.String(); body != expectedRootResponse {
		t.Errorf("Unexpected response body for root path: got %v want %v", body, expectedRootResponse)
	}

	reqAPI, err := http.NewRequest("GET", server.URL+"/api/", nil)
	if err != nil {
		t.Fatal(err)
	}

	wAPI := httptest.NewRecorder()

	r.ServeHTTP(wAPI, reqAPI)

	if status := wAPI.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code for API request: got %v want %v", status, http.StatusOK)
	}

	expectedAPIResponse := "Hello from API"
	if body := wAPI.Body.String(); body != expectedAPIResponse {
		t.Errorf("Unexpected response body for API path: got %v want %v", body, expectedAPIResponse)
	}
}
