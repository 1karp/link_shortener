package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var storage = make(map[string]string, 100)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randSeqGen() string {
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

	responseData, err := io.ReadAll(r.Body)
	if err != nil || string(responseData) == "" {
		http.Error(w, "Invalid POST body!", http.StatusBadRequest)
		return
	}
	url := string(responseData)

	shortURL := randSeqGen()
	storage[shortURL] = url

	w.WriteHeader(http.StatusCreated)
	_, errWrite := w.Write([]byte("http://" + r.Host + "/" + shortURL))
	if errWrite != nil {
		panic(errWrite)
	}

}

func shortenedHandler(w http.ResponseWriter, r *http.Request) {
	d := chi.URLParam(r, "id")

	if full, ok := storage[d]; ok {
		w.Header().Add("Location", full)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Invalid request", http.StatusBadRequest)
}

func main() {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/", mainHandler)
		r.Get("/{id}", shortenedHandler)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
