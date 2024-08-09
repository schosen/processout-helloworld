package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
	errorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP errors",
		},
		[]string{"path"},
	)
)

func init() {
	// Register metrics
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(responseDuration)
	prometheus.MustRegister(errorsTotal)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Job string `json:"job"`
}

var users = []User{
    {ID: 1, Name: "Sarah Chosen", Job: "DevOps Engineer"},
    {ID: 2, Name: "John Doe", Job: "Builder"},
    {ID: 3, Name: "Jane Doe", Job: "Actress"},
}


func main() {
	http.HandleFunc("/", loggingMiddleware(getHelloHandler))
	http.HandleFunc("/users", loggingMiddleware(getUsersHandler))
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func getHelloHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	_, err := fmt.Fprintln(w, "Hello, World!")
	duration := time.Since(start).Seconds()

	// Update Prometheus metrics
	requestsTotal.WithLabelValues("/").Inc()
	responseDuration.WithLabelValues("/").Observe(duration)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		errorsTotal.WithLabelValues("/").Inc()
	}
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	duration := time.Since(start).Seconds()

	// Update Prometheus metrics
	requestsTotal.WithLabelValues("/users").Inc()
	responseDuration.WithLabelValues("/users").Observe(duration)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		errorsTotal.WithLabelValues("/users").Inc()
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
