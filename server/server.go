package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thomas-chastaingt/Goflix/store"
)

type Server struct {
	Router *mux.Router
	Store  store.Store
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}

	return
}

func (s *Server) ServHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
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
