package routes

import (
	"fmt"
	"net/http"
)

func handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix")
	}
}
