package server

import (
	"net/http"
	"time"

	"github.com/NavneetSinghGour/devops-dashboard/internal/config"
	"github.com/NavneetSinghGour/devops-dashboard/internal/router"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg config.Config) *Server {

	router.RegisterRoutes()

	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      nil,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
