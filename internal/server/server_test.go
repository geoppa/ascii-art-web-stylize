package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// TestStart_RoutesRegistration ελέγχει ότι ο router της εφαρμογής
// περιέχει όλα τα routes που πρέπει να είναι διαθέσιμα.
func TestStart_RoutesRegistration(t *testing.T) {

	// Δημιουργούμε τον router της εφαρμογής χωρίς να ξεκινήσουμε
	// πραγματικό HTTP server.
	router := NewRouter()

	tests := []struct {
		name string
		path string
	}{
		{"Home Route Registration", "/"},
		{"AsciiArt Route Registration", "/ascii-art"},
		{"Static Files Route Registration", "/static/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Δημιουργούμε ένα δοκιμαστικό request και ζητάμε από τον router
			// να μας επιστρέψει το route pattern που αντιστοιχεί στο URL.
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			_, pattern := router.Handler(req)

			// Αν το pattern είναι κενό, σημαίνει ότι το route δεν καταχωρήθηκε
			if pattern == "" {
				t.Errorf("Route pattern %s was not properly registered in router", tt.path)
			}
		})
	}
}

// TestStart_StaticCSS ελέγχει αν το αρχείο style.css σερβίρεται σωστά
// και με το κατάλληλο Content-Type header.
func TestStart_StaticCSS(t *testing.T) {
	// Αλλάζουμε το working directory στο root για να βρει τον φάκελο static/
	oldWd, _ := os.Getwd()
	err := os.Chdir("../..")
	if err != nil {
		t.Fatalf("Failed to change working directory to root: %v", err)
	}
	defer os.Chdir(oldWd)

	// Προσομοιώνουμε ένα HTTP GET request για το αρχείο style.css
	req := httptest.NewRequest(http.MethodGet, "/static/style.css", nil)
	rr := httptest.NewRecorder()

	// Δημιουργούμε τον router και προσομοιώνουμε ένα request
	// προς το static αρχείο CSS χωρίς να ανοίξουμε πραγματικό port.
	router := NewRouter()

	// Στέλνουμε το request στον router και καταγράφουμε την απάντηση
	// μέσα στον ResponseRecorder.
	router.ServeHTTP(rr, req)

	// 1. Ελέγχουμε αν το αρχείο σερβίρεται επιτυχώς (200 OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK for style.css, got %d. Make sure 'static/style.css' exists.", rr.Code)
	}

	// 2. Ελέγχουμε αν το Content-Type header είναι σωστά ρυθμισμένο σε text/css
	contentType := rr.Header().Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/css") {
		t.Errorf("Expected Content-Type to start with 'text/css', got '%s'", contentType)
	}
}