package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NavneetSinghGour/devops-dashboard/internal/config"
	"github.com/NavneetSinghGour/devops-dashboard/internal/services"
)

func ApplicationInfo(w http.ResponseWriter, r *http.Request) {

	cfg := config.Load()

	data := services.GetDashboardData(cfg)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}
