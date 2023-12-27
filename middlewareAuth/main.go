package main

import (
	"fmt"
	"net/http"
	"time"
)

func MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Processing your request: %s %s\n", r.Method, r.URL.Path)

		next(w, r)

		duration := time.Since(start)
		fmt.Printf("Request completed in time: %v\n", duration)
	}
}

func MiddlewareAuth(expectedToken string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != expectedToken {
			http.Error(w, "No authorization", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func main() {
	router := http.NewServeMux()

	authenticatedHandler := MiddlewareAuth("Authorization", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	router.HandleFunc("/", MiddlewareLogger(authenticatedHandler))

	http.ListenAndServe(":8080", router)
}
