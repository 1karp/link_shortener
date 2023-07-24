package handlers

import (
	"net/http"

	"github.com/1karp/link_shortener/internal/app/storage"
	"github.com/go-chi/chi/v5"
)

func ShortenedHandler(w http.ResponseWriter, r *http.Request, s storage.Storage) {
	shortURL := chi.URLParam(r, "id")

	if fullURL := s.Get(shortURL); fullURL != "" {
		w.Header().Add("Location", fullURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Invalid request", http.StatusBadRequest)
}
