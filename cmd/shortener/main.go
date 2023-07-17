package main

import (
	"log"

	"github.com/1karp/go-musthave-shortener-tpl/internal/app/config"
	"github.com/1karp/go-musthave-shortener-tpl/internal/app/router"
	"github.com/1karp/go-musthave-shortener-tpl/internal/app/server"
)

func main() {
	cfg := config.NewConfig()

	router := router.NewRouter(cfg)

	log.Printf("Starting server on %s\n", cfg.GetAddress())

	server := server.NewServer(cfg, router.Router)

	err := server.Start()
	if err != nil {
		panic(err)
	}
}
