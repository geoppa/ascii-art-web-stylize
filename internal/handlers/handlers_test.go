package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// setupMockTemplates creates temporary template files required by the handlers
func setupMockTemplates(t *testing.T) func() {
	tmpDir := "templates"
	err := os.MkdirAll(tmpDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp templates directory: %v", err)
	}

	// Minimal valid HTML templates that include the expected Go template pipelines
	indexHTML := `<html><body>{{if .Error}}<div>{{.Error}}</div>{{end}}<div>{{.Result}}</div></body></html>`
	errorHTML := `<html><body><h1>{{.Code}}</h1><p>{{.Message}}</p></body></html>`

	err = os.WriteFile(filepath.Join(tmpDir, "index.html"), []byte(indexHTML), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock index.html: %v", err)
	}

	err = os.WriteFile(filepath.Join(tmpDir, "error.html"), []byte(errorHTML), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock error.html: %v", err)
	}

	// Return a cleanup function to delete the temporary folder after tests run
	return func() {
		os.RemoveAll(tmpDir)
	}
}

func TestHomeHandler(t *testing.T) {
	cleanup := setupMockTemplates(t)
	defer cleanup()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "Valid GET request to root",
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid method POST to root",
			method:         http.MethodPost,
			path:           "/",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid path 404 error",
			method:         http.MethodGet,
			path:           "/invalid-path",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP request
			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create a ResponseRecorder to record the handler's response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HomeHandler)

			// Execute the handler
			handler.ServeHTTP(rr, req)

			// Verify the response status code
			if rr.Code != tt.expectedStatus {
				t.Errorf("HomeHandler returned wrong status: got %v want %v", rr.Code, tt.expectedStatus)
			}
		})
	}
}

func TestAsciiArtHandler_MethodNotAllowed(t *testing.T) {
	cleanup := setupMockTemplates(t)
	defer cleanup()

	// AsciiArtHandler only allows POST requests, testing GET here
	req, err := http.NewRequest(http.MethodGet, "/ascii-art", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AsciiArtHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("AsciiArtHandler standard check failed: got %v want %v", rr.Code, http.StatusMethodNotAllowed)
	}
}

func TestAsciiArtHandler_FormSubmission(t *testing.T) {
	cleanup := setupMockTemplates(t)
	defer cleanup()

	// Set up temporary banners directory since AsciiArtHandler loads banner files internally
	bannerDir := "banners"
	err := os.MkdirAll(bannerDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp banners directory: %v", err)
	}
	defer os.RemoveAll(bannerDir)

	// Create a dummy standard banner file (e.g., standard.txt) to avoid 500 Internal Server Error
	// Note: Dependending on your internal validation/render logic, you might need to structure this content
	err = os.WriteFile(filepath.Join(bannerDir, "standard.txt"), []byte("mock banner data"), 0644)
	if err != nil {
		t.Fatalf("Failed to create mock banner file: %v", err)
	}

	// Prepare form data parameters
	formData := url.Values{}
	formData.Set("text", "Hello")
	formData.Set("banner", "standard")

	req, err := http.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	// Content-Type header is required for ParseForm() to work properly
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AsciiArtHandler)
	handler.ServeHTTP(rr, req)

	// We check if the request completes without system crash.
	// If validation fails inside your custom internal validation package, it returns 400 Bad Request.
	// If everything succeeds, it returns 200 OK.
	if rr.Code != http.StatusOK && rr.Code != http.StatusBadRequest {
		t.Errorf("AsciiArtHandler returned unexpected status code: %v", rr.Code)
	}
}
