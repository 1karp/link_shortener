package server

import (
	"net/http"

	"github.com/1karp/link_shortener/internal/app/config"
	"github.com/1karp/link_shortener/internal/app/handlers"
	"github.com/1karp/link_shortener/internal/app/storage"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	config  config.Config
	router  *chi.Mux
	storage storage.Storage
}

func NewServer(config config.Config, storage storage.Storage) *Server {
	chiRouter := chi.NewRouter()

	router := &Server{
		config:  config,
		router:  chiRouter,
		storage: storage,
	}

	chiRouter.Route("/", func(r chi.Router) {
		r.Post("/", router.MainHandler)
		r.Get("/{id}", router.ShortenedHandler)
	})

	return router
}

func (s *Server) MainHandler(rw http.ResponseWriter, req *http.Request) {
	handlers.MainHandler(rw, req, s.storage)
}

func (s *Server) ShortenedHandler(rw http.ResponseWriter, req *http.Request) {
	handlers.ShortenedHandler(rw, req, s.storage)
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.config.GetAddress(), s.router)
}
