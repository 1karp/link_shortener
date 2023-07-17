package server

import (
	"net/http"

	"github.com/1karp/go-musthave-shortener-tpl/internal/app/config"
)

type Server struct {
	config config.Config
	router http.Handler
}

func NewServer(config config.Config, router http.Handler) *Server {
	return &Server{
		config: config,
		router: router,
	}
}

func (srv *Server) Start() error {
	return http.ListenAndServe(srv.config.GetAddress(), srv.router)
}
