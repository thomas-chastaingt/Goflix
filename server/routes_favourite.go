package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thomas-chastaingt/Goflix/favourite"
)

func (s *Server) handleAddFavourite() http.HandlerFunc {
	type request struct {
		idUser  int `json:"idUser"`
		idMovie int `json:"idMovie"`
	}
	type response struct{}
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
		fmt.Println("id movie : ", req.idMovie)

		f := &favourite.Favourite{
			IDMovie: req.idMovie,
			IDUser:  req.idUser,
		}

		err = s.Store.CreateFavourite(f)
		if err != nil {
			log.Printf("Cannot add movie to favourite")
			s.Respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		s.Respond(w, r, response{}, http.StatusOK)

	}
}
