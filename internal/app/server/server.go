package server

import (
	"log"
	"net/http"
	"os"

	"github.com/1karp/link_shortener/internal/app/config"
	"github.com/1karp/link_shortener/internal/app/handlers"
	"github.com/1karp/link_shortener/internal/app/logging"
	"github.com/1karp/link_shortener/internal/app/storage"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Server struct {
	config  config.Config
	router  *chi.Mux
	storage storage.Storage
}

func NewServer(config config.Config, storage storage.Storage) *Server {
	chiRouter := chi.NewRouter()

	server := &Server{
		config:  config,
		router:  chiRouter,
		storage: storage,
	}

	logger, err := zap.NewDevelopment()

	if err != nil {
		log.Fatal("Error initializing logger", err)
		os.Exit(1)
	}

	logging.Sugar = logger.Sugar()
	chiRouter.Use(logging.CustomMiddlewareLogger)
	chiRouter.Route("/", func(r chi.Router) {
		r.Post("/", server.MainHandler)
		r.Get("/{id}", server.ShortenedHandler)
		r.Post("/api/shorten", http.HandlerFunc(server.APIShortenHandler))
	})

	return server
}

func (s *Server) MainHandler(rw http.ResponseWriter, req *http.Request) {
	handlers.MainHandler(rw, req, s.storage)
}

func (s *Server) ShortenedHandler(rw http.ResponseWriter, req *http.Request) {
	handlers.ShortenedHandler(rw, req, s.storage)
}

func (s *Server) APIShortenHandler(rw http.ResponseWriter, req *http.Request) {
	handlers.APIShortenHandler(rw, req, s.storage, s.config)
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.config.GetAddress(), s.router)
}
