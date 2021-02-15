package server

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix")
	}
}

func (s *Server) handleTokenCreate() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type response struct {
		Token string `json:"token"`
	}
	type respondError struct {
		Error string `json:"error"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix")
		req := request{}
		err := s.decode(w, r, &req)
		if err != nil {
			msg := fmt.Sprintf("Cannot parse login body. err=%v", err)
			log.Println(msg)
			s.Respond(w, r responseError{
				Error: msg,
			}, http.StatusBadRequest)
			return
		}
		if req.Username != "golang" 
	}
}
