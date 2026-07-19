package server

import (
	"net/http"
	"time"

	"github.com/NavneetSinghGour/devops-dashboard/internal/config"
	"github.com/NavneetSinghGour/devops-dashboard/internal/router"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg config.Config) *Server {

	router.RegisterRoutes()

	return &Server{
		httpServer: &http.Server{
			Addr: ":" + cfg.Port,
			Handler: otelhttp.NewHandler(
				http.DefaultServeMux,
				"http-server",
			),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
