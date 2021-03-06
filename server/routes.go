package server

//Routes define all routes
func (s *Server) Routes() {
	s.Router.HandleFunc("/", s.handleIndex()).Methods("GET")
	//Routes favourites
	s.Router.HandleFunc("/api/favourite", s.handleAddFavourite()).Methods("POST")
	//Routes movie
	s.Router.HandleFunc("/api/movies", s.middleware(s.handleMovieList())).Methods("GET")
	s.Router.HandleFunc("/api/movies", s.handleMovieCreate()).Methods("POST")
	s.Router.HandleFunc("/api/movies/{id:[0-9]+}", s.handleMovieDetail()).Methods("GET")
	//Routes user
	s.Router.HandleFunc("/api/user/login", s.handleUserLogin()).Methods("POST")
	s.Router.HandleFunc("/api/user/create", s.handleUserCreate()).Methods("POST")
}
