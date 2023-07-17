package main

import (
	"log"
	"net/http"

	"github.com/1karp/go-musthave-shortener-tpl/internal/app/config"
	"github.com/1karp/go-musthave-shortener-tpl/internal/app/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.NewConfig()

	r := chi.NewRouter()
	r.Post("/", handlers.MainHandler)
	r.Get("/{id}", handlers.ShortenedHandler)

	log.Printf("Starting server on %s\n", cfg.GetAddress())

	log.Fatal(http.ListenAndServe(cfg.GetAddress(), r))
}
