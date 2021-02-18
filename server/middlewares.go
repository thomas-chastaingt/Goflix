package server

import (
	"log"
	"net/http"
)

func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v", r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	}
}

// func (s *Server) loggedOnly(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		j := jwtmiddleware.New(jwtmiddleware.Options{
// 			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 				return []byte(JWT_APP_KEY), nil
// 			},
// 			SigningMethod: jwt.SigningMethodHS256,
// 		})
// 		j.HandlerWithNext(w, r, next)
// 	}
// }
