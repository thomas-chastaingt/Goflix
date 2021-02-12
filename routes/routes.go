package routes

import "github.com/thomas-chastaingt/Goflix/server"

func (s *server.Server) routes() {
	s.Router.HandleFunc("/", s.HandleIndex).Methods("GET")
}
