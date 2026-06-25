package handlers

import (
	"html/template"
	"net/http"

	"ascii-art-web/internal/banner"
	"ascii-art-web/internal/render"
	"ascii-art-web/internal/validation"
)

// PageData holds form and result information for the main UI
type PageData struct {
	Text   string
	Banner string
	Error  string
	Result string
}

// ErrorPageData holds values to display custom error screens
type ErrorPageData struct {
	Code    int
	Message string
}

// renderHomePage displays the primary user submission layout
func renderHomePage(w http.ResponseWriter, data PageData) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// If the main template is missing, it is a server error (500), not a 404
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

// renderErrorPage outputs dedicated custom error HTML screens
func renderErrorPage(w http.ResponseWriter, statusCode int, message string) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// Fallback to plain text if the error template file is missing
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Explicitly set the response status code header before rendering
	w.WriteHeader(statusCode)

	err = tmpl.Execute(w, ErrorPageData{
		Code:    statusCode,
		Message: message,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HomeHandler serves the dashboard interface and protects against broken paths
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Reject paths that do not exactly match the home domain root
	if r.URL.Path != "/" {
		renderErrorPage(w, http.StatusNotFound, "Page Not Found")
		return
	}

	// Reject any incoming method that is not a standard GET request
	if r.Method != http.MethodGet {
		renderErrorPage(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	renderHomePage(w, PageData{})
}

// AsciiArtHandler validates form inputs, reads files, and processes the output art
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Block requests trying to execute anything other than a secure POST
	if r.Method != http.MethodPost {
		renderErrorPage(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Read form inputs safely
	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, http.StatusBadRequest, "Bad Request")
		return
	}

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")

	// Verify inputted text conforms to printable character rules
	err = validation.ValidateText(text)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // tells the browser this is a 400 error
		renderHomePage(w, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  err.Error(),
		})
		return
	}

	// Verify the chosen template target name matches existing files
	err = validation.ValidateBanner(bannerName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // force a 400 Bad Request network header
		renderHomePage(w, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  err.Error(),
		})
		return
	}

	bannerFile := bannerName + ".txt"

	// Read text representation matrix out of storage
	bannerData, err := banner.Load(bannerFile)
	if err != nil {
		renderErrorPage(
			w,
			http.StatusInternalServerError,
			"Internal Server Error",
		)
		return
	}

	// Generate standard visual ASCII character block map
	result := render.Generate(text, bannerData)

	// Send execution success back to UI output target elements
	renderHomePage(w, PageData{
		Text:   text,
		Banner: bannerName,
		Result: result,
	})
}
