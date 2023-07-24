package handlers

import (
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/1karp/link_shortener/internal/app/storage"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateShortURL() string {
	shortURL := make([]byte, 8)

	for i := range shortURL {
		shortURL[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(shortURL)
}

func MainHandler(w http.ResponseWriter, r *http.Request, s storage.Storage) {
	body, err := io.ReadAll(r.Body)
	if err != nil || string(body) == "" {
		http.Error(w, "Invalid POST body!", http.StatusBadRequest)
		return
	}

	url := string(body)
	shortURL := generateShortURL()
	s.Set(shortURL, url)

	w.WriteHeader(http.StatusCreated)
	_, errWrite := w.Write([]byte("http://" + r.Host + "/" + shortURL))
	if errWrite != nil {
		log.Printf("Error writing response: %v", errWrite)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
