package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

// TestAsciiArtHandler ελέγχει την παραγωγή του art (POST, 400 για λάθος inputs, 405 για GET)
func TestAsciiArtHandler(t *testing.T) {
	oldWd, _ := os.Getwd()
	err := os.Chdir("../..")
	if err != nil {
		t.Fatalf("Failed to change directory to root: %v", err)
	}
	defer os.Chdir(oldWd)

	// Case 1: Ένα έγκυρο POST request με σωστά Form Data πρέπει να επιστρέψει 200 OK
	t.Run("Valid POST Request", func(t *testing.T) {
		form := url.Values{}
		form.Add("text", "Hello")
		form.Add("banner", "standard")

		req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		AsciiArtHandler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("AsciiArtHandler returned wrong status: got %d want %d", rr.Code, http.StatusOK)
		}
	})

	// Case 2: Ένα POST request με μη-ASCII χαρακτήρες πρέπει να επιστρέψει 400 Bad Request
	t.Run("Invalid Characters 400", func(t *testing.T) {
		form := url.Values{}
		form.Add("text", "Καλημέρα") // Ελληνικά (Μη-ASCII)
		form.Add("banner", "standard")

		req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		AsciiArtHandler(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("AsciiArtHandler returned wrong status for Greek text: got %d want %d", rr.Code, http.StatusBadRequest)
		}
	})

	// Case 3: Μη έγκυρο banner πρέπει να επιστρέψει 400 Bad Request
	t.Run("Invalid Banner 400", func(t *testing.T) {
		form := url.Values{}
		form.Add("text", "Hello")
		form.Add("banner", "random")

		req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		AsciiArtHandler(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("AsciiArtHandler returned wrong status for invalid banner: got %d want %d", rr.Code, http.StatusBadRequest)
		}
	})

// Case 4: Κενό κείμενο πρέπει να επιστρέψει 400 Bad Request
t.Run("Empty Text 400", func(t *testing.T) {
	form := url.Values{}
	form.Add("text", "")
	form.Add("banner", "standard")

	req := httptest.NewRequest(
		http.MethodPost,
		"/ascii-art",
		strings.NewReader(form.Encode()),
	)

	req.Header.Add(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	rr := httptest.NewRecorder()

	AsciiArtHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf(
			"AsciiArtHandler returned wrong status for empty text: got %d want %d",
			rr.Code,
			http.StatusBadRequest,
		)
	}
})

	// Case 5: Ένα GET request στο "/ascii-art" πρέπει να αποκλειστεί με 405 Method Not Allowed
	t.Run("Invalid GET Method 405", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
		rr := httptest.NewRecorder()

		AsciiArtHandler(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("AsciiArtHandler returned wrong status for GET: got %d want %d", rr.Code, http.StatusMethodNotAllowed)
		}
	})
}