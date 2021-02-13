package routes

import (
	"fmt"
	"net/http"

	"github.com/thomas-chastaingt/Goflix/server"
)

func (s *server.Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix")
	}
}
