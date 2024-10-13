package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Suryavanshi98/url-shortener/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the URL Shortener!")
	})

	http.HandleFunc("/shorten", handlers.ShortenURLHandler)
	http.HandleFunc("/s/", handlers.RedirectHandler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
