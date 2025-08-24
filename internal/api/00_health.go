package api

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := HealthCheckResponse{
		Status: "OK",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

func CatchAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := ErrorResponse{
		Error: "endpoint not found",
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(body)
}
