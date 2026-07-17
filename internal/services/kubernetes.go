package services

import (
	"os"

	"github.com/NavneetSinghGour/devops-dashboard/internal/models"
)

func GetKubernetesInfo() *models.KubernetesInfo {

	return &models.KubernetesInfo{
		PodName:   getEnv("POD_NAME"),
		Namespace: getEnv("POD_NAMESPACE"),
		NodeName:  getEnv("NODE_NAME"),
		PodIP:     getEnv("POD_IP"),
		HostIP:    getEnv("HOST_IP"),
	}
}

func getEnv(key string) string {

	value := os.Getenv(key)

	if value == "" {
		return "--"
	}

	return value
}
