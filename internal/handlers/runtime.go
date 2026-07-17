package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NavneetSinghGour/devops-dashboard/internal/services"
)

func RuntimeInfo(w http.ResponseWriter, r *http.Request) {

	info := services.GetRuntimeInfo()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(info)
}
