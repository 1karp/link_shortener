package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	d := strings.TrimPrefix(r.URL.Path, "/")

	if d == "" {
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

		shortUrl := randSeqGen()
		storage[shortUrl] = url

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("http://" + r.Host + "/" + shortUrl))
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", http.StatusBadRequest)
		return
	}

	if full, ok := storage[d]; ok {
		w.Header().Add("Location", full)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Invalid request", http.StatusBadRequest)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
