package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thomas-chastaingt/Goflix/movie"
)

type JsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerURL  string `json:"trailer_url"`
}

func (s *Server) handleMovieList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := s.Store.GetMovies()
		if err != nil {
			log.Printf("Cannot load movies, err=%v\n", err)
			s.Respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		var resp = make([]JsonMovie, len(movies))
		for i, m := range movies {
			resp[i] = mapMovieToJson(m)
		}

		s.Respond(w, r, resp, http.StatusOK)
	}
}

func (s *Server) handleMovieDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Printf("Cannot parse id to int")
			s.Respond(w, r, nil, http.StatusBadRequest)
			return
		}
		m, err := s.Store.GetMovieById(id)
		if err != nil {
			log.Printf("Cannot load movie. err=%v", err)
			s.Respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		var resp = mapMovieToJson(m)
		s.Respond(w, r, resp, http.StatusOK)
	}
}

func (s *Server) handleMovieCreate() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Duration    int    `json:"duration"`
		TrailerURL  string `json:"trailer_url"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := s.decode(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse movie body error = %v", err)
			s.Respond(w, r, nil, http.StatusBadRequest)
			return
		}
		m := &movie.Movie{
			ID:          0,
			Title:       req.Title,
			ReleaseDate: req.ReleaseDate,
			Duration:    req.Duration,
			TrailerURL:  req.TrailerURL,
		}

		err = s.Store.CreateMovie(m)
		if err != nil {
			log.Printf("Cannot create a movie")
			s.Respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		var resp = mapMovieToJson(m)
		s.Respond(w, r, resp, http.StatusOK)

	}
}

func mapMovieToJson(m *movie.Movie) JsonMovie {
	return JsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerURL:  m.TrailerURL,
	}
}
