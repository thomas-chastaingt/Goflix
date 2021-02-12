package routes

import (
	"github.com/thomas-chastaingt/Goflix/server"
)

func (s *server.Server) routes() {
	s.router.HandleFunc("/", nil).Methods("GET")
}
