package router

import (
	"github.com/1karp/go-musthave-shortener-tpl/internal/app/config"
	"github.com/1karp/go-musthave-shortener-tpl/internal/app/handlers"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	Router *chi.Mux
	config config.Config
}

func NewRouter(config config.Config) *Router {
	chiRouter := chi.NewRouter()

	router := &Router{
		Router: chiRouter,
		config: config,
	}

	chiRouter.Route("/", func(r chi.Router) {
		r.Post("/", handlers.MainHandler)
		r.Get("/{id}", handlers.ShortenedHandler)
	})

	return router
}
