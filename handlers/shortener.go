package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"
)

var (
	urlStore = make(map[string]string)
	mu       sync.Mutex
)

// ShortenURLHandler handles shortening URLs
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL(originalURL)

	mu.Lock()
	urlStore[shortURL] = originalURL
	mu.Unlock()

	shortened := fmt.Sprintf("http://localhost:8080/s/%s", shortURL)
	fmt.Fprintf(w, "Shortened URL: %s", shortened)
}

// Generate a simple SHA-1 hash of the original URL to create a short version
func generateShortURL(originalURL string) string {
	hash := sha1.New()
	hash.Write([]byte(originalURL))
	shortHash := hex.EncodeToString(hash.Sum(nil))[:6]
	return shortHash
}

// RedirectHandler handles redirection to the original URL
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL from the path
	shortURL := r.URL.Path[len("/s/"):]

	mu.Lock()
	originalURL, exists := urlStore[shortURL]
	mu.Unlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}
