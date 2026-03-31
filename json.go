package main

import (
    "encoding/json"
    "net/http"
)

type StatusResponse struct {
    Message string `json:"message"`
    Status  string `json:"status"`
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    response := StatusResponse{
        Message: "Server is healthy",
        Status:  "OK",
    }

    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/health", healthcheck)
    http.ListenAndServe(":8080", nil)
}
