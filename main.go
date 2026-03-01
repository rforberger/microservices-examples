package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Service   string `json:"service"`
	Version   string `json:"version"`
	Hostname  string `json:"hostname"`
	Timestamp string `json:"timestamp"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	service := getEnv("SERVICE_NAME", "info-service")
	version := getEnv("VERSION", "dev")
	hostname, _ := os.Hostname()

	resp := Response{
		Service:   service,
		Version:   version,
		Hostname:  hostname,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}