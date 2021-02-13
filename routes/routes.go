package routes

import "github.com/thomas-chastaingt/Goflix/server"

func (s *server.Server) Routes() {
	s.Router.HandleFunc("/", s.HandleIndex).Methods("GET")
	s.Router.HandleFunc("/api/movies/", s.handleMovieList()).Methods("GET")
}
