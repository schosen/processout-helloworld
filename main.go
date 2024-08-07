package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define metrics
var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
	responseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_duration_seconds",
			Help:    "Duration of HTTP responses",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func init() {
	// Register metrics
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(responseDuration)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Increment request counter
	requestsTotal.WithLabelValues(r.URL.Path).Inc()

	// Write response
	fmt.Fprintln(w, "Hello, World!")

	// Observe response duration
	duration := time.Since(start).Seconds()
	responseDuration.WithLabelValues(r.URL.Path).Observe(duration)
}

func main() {
	// Register handlers
	http.HandleFunc("/", helloHandler)
	http.Handle("/metrics", promhttp.Handler())

	// Start server
	port := 8080
	log.Printf("Starting server on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
