package main

import (
	"log"

	"github.com/1karp/link_shortener/internal/app/config"
	"github.com/1karp/link_shortener/internal/app/server"
	"github.com/1karp/link_shortener/internal/app/storage"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("No config found: %v", err)
	}

	stor := storage.NewStorage()

	log.Printf("Starting server on %s\n", cfg.GetAddress())
	server := server.NewServer(cfg, stor)

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
