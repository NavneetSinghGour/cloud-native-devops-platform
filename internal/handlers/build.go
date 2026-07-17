package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NavneetSinghGour/devops-dashboard/internal/services"
)

func BuildInfo(w http.ResponseWriter, r *http.Request) {

	build := services.GetBuildInfo()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(build)
}
