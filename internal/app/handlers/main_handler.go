package handlers

import (
	"io"
	"math/rand"
	"net/http"

	"github.com/1karp/go-musthave-shortener-tpl/internal/app/storage"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateShortURL() string {
	shortURL := make([]byte, 8)

	for i := range shortURL {
		shortURL[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(shortURL)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed!", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || string(body) == "" {
		http.Error(w, "Invalid POST body!", http.StatusBadRequest)
		return
	}

	url := string(body)
	shortURL := generateShortURL()
	storage.Set(shortURL, url)

	w.WriteHeader(http.StatusCreated)
	_, errWrite := w.Write([]byte("http://" + r.Host + "/" + shortURL))
	if errWrite != nil {
		panic(errWrite)
	}
}
