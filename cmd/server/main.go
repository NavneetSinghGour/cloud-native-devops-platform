package main

import (
	"github.com/NavneetSinghGour/devops-dashboard/internal/config"
	"github.com/NavneetSinghGour/devops-dashboard/internal/logger"
	"github.com/NavneetSinghGour/devops-dashboard/internal/server"
)

func main() {

	cfg := config.Load()

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
