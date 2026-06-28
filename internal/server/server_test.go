package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestNewRouter(t *testing.T) {
	// Create a temporary static folder and a sample CSS file to test the static file server route
	staticDir := "static"
	err := os.MkdirAll(staticDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temporary static directory: %v", err)
	}
	defer os.RemoveAll(staticDir)

	mockCSSContent := "body { background: #fff; }"
	err = os.WriteFile(filepath.Join(staticDir, "style.css"), []byte(mockCSSContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock asset file: %v", err)
	}

	// Create temporary templates directory since handlers mapped to the router depend on them
	tmpTemplatesDir := "templates"
	err = os.MkdirAll(tmpTemplatesDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp templates directory: %v", err)
	}
	defer os.RemoveAll(tmpTemplatesDir)

	// Setup basic HTML files so handlers can render properly without returning 500
	err = os.WriteFile(filepath.Join(tmpTemplatesDir, "index.html"), []byte("<html></html>"), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock index template: %v", err)
	}
	err = os.WriteFile(filepath.Join(tmpTemplatesDir, "error.html"), []byte("<html></html>"), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock error template: %v", err)
	}

	// Create temporary banners directory to satisfy internal handler dependencies
	bannerDir := "banners"
	err = os.MkdirAll(bannerDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp banners directory: %v", err)
	}
	defer os.RemoveAll(bannerDir)

	// Initialize the router under test
	router := NewRouter()

	tests := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
	}{
		{
			name:           "Root path maps to HomeHandler",
			method:         http.MethodGet,
			url:            "/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Ascii-art path maps to AsciiArtHandler (Method Not Allowed for GET)",
			method:         http.MethodGet,
			url:            "/ascii-art",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Static file asset distribution check",
			method:         http.MethodGet,
			url:            "/static/style.css",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Static missing asset distribution check",
			method:         http.MethodGet,
			url:            "/static/missing.css",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatalf("Failed to build request instance: %v", err)
			}

			rr := httptest.NewRecorder()

			// ServeHTTP executes the matching handler registered inside the custom ServeMux
			router.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Router path %s returned unexpected status: got %v want %v", tt.url, rr.Code, tt.expectedStatus)
			}
		})
	}
}
