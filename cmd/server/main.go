package main

import (
	"context"

	"github.com/NavneetSinghGour/devops-dashboard/internal/config"
	"github.com/NavneetSinghGour/devops-dashboard/internal/logger"
	"github.com/NavneetSinghGour/devops-dashboard/internal/observability"
	"github.com/NavneetSinghGour/devops-dashboard/internal/server"
)

func main() {

	cfg := config.Load()

	// Initialize OpenTelemetry
	tp, err := observability.InitTracer(cfg.AppName)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Error(err.Error())
		}
	}()

	srv := server.New(cfg)

	logger.Info("====================================")
	logger.Info("Starting " + cfg.AppName)
	logger.Info("Version: " + cfg.Version)
	logger.Info("Environment: " + cfg.Environment)
	logger.Info("Listening on :" + cfg.Port)
	logger.Info("====================================")

	if err := srv.Start(); err != nil {
		logger.Fatal(err.Error())
	}
}
