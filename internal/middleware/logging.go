package middleware

import (
	"net/http"

	"github.com/NavneetSinghGour/devops-dashboard/internal/logger"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger.Info(r.Method + " " + r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
