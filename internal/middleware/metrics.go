package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/NavneetSinghGour/devops-dashboard/internal/metrics"
)

type MetricsResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *MetricsResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Metrics(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		metrics.InFlightRequests.Inc()

		rw := &MetricsResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()

		metrics.RequestCount.WithLabelValues(
			r.Method,
			r.URL.Path,
			strconv.Itoa(rw.statusCode),
		).Inc()

		metrics.RequestDuration.WithLabelValues(
			r.Method,
			r.URL.Path,
		).Observe(duration)

		metrics.InFlightRequests.Dec()
	})
}
