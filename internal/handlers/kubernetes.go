package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NavneetSinghGour/devops-dashboard/internal/services"
)

func KubernetesInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(services.GetKubernetesInfo())
}
