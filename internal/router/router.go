package router

import (
	"net/http"

	"github.com/NavneetSinghGour/devops-dashboard/internal/handlers"
	"github.com/NavneetSinghGour/devops-dashboard/internal/middleware"
)

func RegisterRoutes() {

	// Create a new router
	mux := http.NewServeMux()

	// Static Files
	fs := http.FileServer(http.Dir("internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Dashboard
	mux.HandleFunc("/", handlers.Dashboard)

	// Health
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/ready", handlers.Health)
	mux.HandleFunc("/live", handlers.Health)

	// APIs
	mux.HandleFunc("/api/info", handlers.ApplicationInfo)
	mux.HandleFunc("/api/build", handlers.BuildInfo)
	mux.HandleFunc("/api/runtime", handlers.RuntimeInfo)
	mux.HandleFunc("/api/kubernetes", handlers.KubernetesInfo)

	// Apply Logging Middleware
	http.Handle("/", middleware.Logging(mux))
}
