package main

import (
	"log"
	"os"

	"github.com/1karp/link_shortener/internal/app/config"
	"github.com/1karp/link_shortener/internal/app/server"
	"github.com/1karp/link_shortener/internal/app/storage"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading config", err)
		os.Exit(1)
	}

	stor := storage.NewStorage()

	log.Printf("Starting server on %s\n", cfg.GetAddress())
	server := server.NewServer(cfg, stor)

	err = server.Start()
	if err != nil {
		log.Fatal("Server start error", err)
		os.Exit(1)
	}
}
