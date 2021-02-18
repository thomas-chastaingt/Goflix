package server

func (s *Server) Routes() {
	s.Router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.Router.HandleFunc("/api/user/login", s.handleUserLogin()).Methods("POST")
	s.Router.HandleFunc("/api/user/create", s.handleUserCreate()).Methods("POST")
	s.Router.HandleFunc("/api/movies", s.handleMovieList()).Methods("GET")
	s.Router.HandleFunc("/api/movies", s.handleMovieCreate()).Methods("POST")
	s.Router.HandleFunc("/api/movies/{id:[0-9]+}", s.handleMovieDetail()).Methods("GET")
}
