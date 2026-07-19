package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (

	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	InFlightRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Current number of active HTTP requests.",
		},
	)
)

func Init() {

	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(InFlightRequests)

}
