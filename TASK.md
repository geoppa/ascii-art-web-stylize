# 📋 Project Task Checklist - ASCII Art Web

This document tracks the implementation progress, completed features, testing coverage, and improvements made throughout the development of the **ASCII Art Web** project.

---

# 🟥 1. Core Server & Routing

- [x] Create the web server using the Go standard library (`net/http`).
- [x] Register the **Home** endpoint (`GET /`).
- [x] Register the **ASCII Art** endpoint (`POST /ascii-art`).
- [x] Serve static assets through the `/static/` route.
- [x] Keep routing logic isolated inside the `server` package.
- [x] Build the application router using `http.NewServeMux()`.

---

# 🟧 2. HTTP Status Code Handling

- [x] Return **200 OK** for successful page rendering.
- [x] Return **400 Bad Request** for invalid user input.
- [x] Return **404 Not Found** for unknown routes.
- [x] Return **405 Method Not Allowed** for unsupported HTTP methods.
- [x] Return **500 Internal Server Error** when templates or banner files cannot be loaded.

---

# 🟨 3. Input Validation & Security

- [x] Reject empty submissions.
- [x] Reject non-printable and non-ASCII characters.
- [x] Accept printable ASCII characters (32–126).
- [x] Support multiline text input.
- [x] Validate banner selection before loading files.
- [x] Reject invalid banner names.

---

# 🟩 4. Banner File Processing

- [x] Load banner files from the `banners/` directory.
- [x] Read banner files line by line.
- [x] Remove Windows carriage returns (`\r`) for cross-platform compatibility.
- [x] Return file system errors when banner resources cannot be opened.

---

# 🟦 5. ASCII Art Rendering

- [x] Convert literal `\\n` sequences into real newlines.
- [x] Normalize Windows line endings (`\r\n`).
- [x] Split multiline input correctly.
- [x] Render printable ASCII characters using the selected banner.
- [x] Preserve empty lines between text blocks.
- [x] Remove the unwanted trailing newline from the final output.

---

# 🟪 6. User Interface

- [x] Build the interface using HTML templates.
- [x] Display generated ASCII art inside a `<pre>` element.
- [x] Preserve submitted text after validation errors.
- [x] Preserve the selected banner after form submission.
- [x] Display validation messages directly on the page.
- [x] Provide a page reset button.
- [x] Display custom error pages for HTTP errors.

---

# 🟫 7. CSS & User Experience (Stylize)

- [x] Create a custom responsive layout.
- [x] Apply a consistent visual design across the application.
- [x] Improve usability with hover and focus effects.
- [x] Provide clear visual feedback for user interactions.
- [x] Ensure generated ASCII art remains readable.
- [x] Support desktop and mobile devices using responsive media queries.

---

# ⬜ 8. Testing

- [x] Integration test for the application router.
- [x] Verify HTTP route registration.
- [x] Test static file serving.
- [x] Test banner file loading.
- [x] Test banner loading failure.
- [x] Test ASCII art generation.
- [x] Test input validation.
- [x] Test successful handler responses.
- [x] Test handler error responses.
- [x] Verify HTTP status code handling.

---

# 🛠 Bug Fixes & Improvements

- [x] Fixed duplicate `WriteHeader()` calls.
- [x] Fixed unwanted trailing newline generation.
- [x] Added custom HTML error pages.
- [x] Preserved form state after validation failures.
- [x] Added responsive CSS styling.
- [x] Added static asset routing.
- [x] Improved Windows compatibility by removing carriage returns from banner files.
- [x] Added automated tests covering the application's core packages.