package server

import (
	"context"
	"net/http"

	"github.com/PeterBooker/summit/internal/config"
	"github.com/PeterBooker/summit/internal/log"
	"github.com/go-chi/chi"
)

// Server holds all the data the App needs
type Server struct {
	router *chi.Mux
	http   *http.Server
	log    *log.Logger
	cfg    *config.Config
}

// New returns a pointer to the main server struct
func New(log *log.Logger, cfg *config.Config) *Server {
	return &Server{
		log: log,
		cfg: cfg,
	}
}

// Setup starts the HTTP Server
func (s *Server) Setup() {
	s.startHTTP()
}

// Shutdown will release resources and stop the server.
func (s *Server) Shutdown(ctx context.Context) {
	s.http.Shutdown(ctx)
}
