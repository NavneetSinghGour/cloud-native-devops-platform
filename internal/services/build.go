package services

import "github.com/NavneetSinghGour/devops-dashboard/internal/models"

func GetBuildInfo() models.BuildInfo {
	return models.BuildInfo{
		Version:   "1.0.0",
		Commit:    "development",
		Branch:    "main",
		BuildTime: "local",
	}
}
