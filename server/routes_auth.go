package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix")
	}
}

func (s *Server) handleUserLogin() http.HandlerFunc {
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
		req := request{}
		err := s.decode(w, r, &req)
		if err != nil {
			msg := fmt.Sprintf("Cannot parse login body. err=%v", err)
			log.Println(msg)
			s.Respond(w, r, respondError{
				Error: msg,
			}, http.StatusBadRequest)
			return
		}

		found, err := s.Store.Fin

		if req.Username != "golang" || req.Password != "rocks" {
			s.Respond(w, r, respondError{
				Error: "Invalid credentials",
			}, http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": req.Username,
			"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			"iat":      time.Now().Unix(),
		})

		tokenStr, err := token.SignedString([]byte(JWT_APP_KEY))
		if err != nil {
			msg := fmt.Sprint("Cannot generate JWT err=%v", err)
			s.Respond(w, r, respondError{
				Error: msg,
			}, http.StatusInternalServerError)
			return
		}
		s.Respond(w, r, response{
			Token: tokenStr,
		}, http.StatusOK)
	}
}
