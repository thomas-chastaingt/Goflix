package server

import (
	"log"
	"net/http"
)

func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v", r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	}
}
