package server

import (
	"github.com/gorilla/mux"
	"github.com/thomas-chastaingt/Goflix/store"
)

type Server struct {
	router *mux.Router
	Store  store.Store
}

func NewServer() *Server {
	s := &Server{}
	return s
}
