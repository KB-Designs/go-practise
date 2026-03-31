package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs details about each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
	//	next.ServeHTTP(w, r) // Call the next handler in the chain
		log.Printf("%s %s %s %s", r.Method, r.RequestURI, time.Since(start), r.RemoteAddr)
	})
}

// HomeHandler is a simple handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Apply the middleware to the HomeHandler
	mux.Handle("/", LoggingMiddleware(http.HandlerFunc(HomeHandler)))

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
