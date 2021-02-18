package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/thomas-chastaingt/Goflix/store"
)

const JWT_APP_KEY = "goflix.go"

type Server struct {
	Router *mux.Router
	Store  store.Store
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"authorization"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)

	s.Router.Use(cors)

	s.Routes()
	return s
}

func (s *Server) ServHTTP(w http.ResponseWriter, r *http.Request) {
	logRequestMiddleware(s.Router.ServeHTTP).ServeHTTP(w, r)
}

func (s *Server) Respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format to json")
	}

}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
