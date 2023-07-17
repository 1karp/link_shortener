package main

import (
	"log"

	"github.com/1karp/go-musthave-shortener-tpl/internal/app/config"
	"github.com/1karp/go-musthave-shortener-tpl/internal/app/router"
	"github.com/1karp/go-musthave-shortener-tpl/internal/app/server"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	router := router.NewRouter(cfg)

	log.Printf("Starting server on %s\n", cfg.Address)

	server := server.NewServer(cfg, router.Router)

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
