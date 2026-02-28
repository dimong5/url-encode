package main

import (
	"fmt"
	"log"
	"net/http"
	"url-encode/internal/service"
	"url-encode/internal/storage"
)

func main() {
	store := storage.NewURLStore()
	svc := service.NewService(store)

	http.HandleFunc("/shorten", handleShorten(svc))
	http.HandleFunc("/redirect/", handleRedirect(svc))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleShorten(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		originalURL := r.FormValue("url")
		if originalURL == "" {
			http.Error(w, "url parameter required", http.StatusBadRequest)
			return
		}

		shortURL := svc.CreateShortURL(originalURL)
		fmt.Fprintf(w, "http://localhost:8080/redirect/%s", shortURL)
	}
}

func handleRedirect(svc *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.PathValue("short")
		if shortURL == "" {
			http.Error(w, "short URL required", http.StatusBadRequest)
			return
		}

		originalURL, err := svc.GetOriginalURL(shortURL)
		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, originalURL, http.StatusFound)
	}
}
