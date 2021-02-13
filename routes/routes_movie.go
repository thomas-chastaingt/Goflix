package route

import (
	"log"
	"net/http"

	"github.com/thomas-chastaingt/Goflix/movie"
	"github.com/thomas-chastaingt/Goflix/server"
)

type JsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerURL  string `json:"trailer_url"`
}

func (s *server.Server) handleMovieList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := s.store.GetMovies()
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

func mapMovieToJson(m *movie.Movie) JsonMovie {
	return JsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerURL:  m.TrailerURL,
	}
}
