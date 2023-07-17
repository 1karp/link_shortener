package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/1karp/go-musthave-shortener-tpl/internal/app/config"
)

var storage = make(map[string]string, 100)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func generateShortURL() string {
	shortURL := make([]byte, 8)

	for i := range shortURL {
		shortURL[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(shortURL)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
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
	storage[shortURL] = url

	w.WriteHeader(http.StatusCreated)
	_, errWrite := w.Write([]byte("http://" + r.Host + "/" + shortURL))
	if errWrite != nil {
		panic(errWrite)
	}
}

func shortenedHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "id")

	if fullURL, ok := storage[shortURL]; ok {
		w.Header().Add("Location", fullURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Invalid request", http.StatusBadRequest)
}

func main() {
	cfg := config.NewConfig()

	r := chi.NewRouter()
	r.Post("/", mainHandler)
	r.Get("/{id}", shortenedHandler)

	log.Printf("Starting server on %s\n", cfg.GetAddress())

	log.Fatal(http.ListenAndServe(cfg.GetAddress(), r))
}
