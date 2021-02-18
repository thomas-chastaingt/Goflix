package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	userAccount "github.com/thomas-chastaingt/Goflix/user"
	"github.com/thomas-chastaingt/Goflix/utils"
)

//JsonUser define User in json
type JsonUser struct {
	ID       int64  `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
}

//handleIndex is the index
func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix")
	}
}

//handleUserLogin permits to sign in
func (s *Server) handleUserLogin() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type response struct {
		Token string   `json:"token"`
		User  JsonUser `json:"user"`
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

		user, err := s.Store.FindUserByName(req.Username)
		if err != nil {
			msg := fmt.Sprint("Cannot find user err=%v", err)
			s.Respond(w, r, respondError{
				Error: msg,
			}, http.StatusInternalServerError)
			return
		}

		if utils.CheckPasswordHash(req.Password, user.Password) == false {
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
		jsonUser := mapUserToJson(user)
		fmt.Println(jsonUser)
		s.Respond(w, r, response{
			Token: tokenStr,
			User:  jsonUser,
		}, http.StatusOK)
	}
}

//handleUserCreate permits to create a new user
func (s *Server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Username       string `json:"username"`
		Password       string `json:"password"`
		VerifyPassword string `json:"verifyPassword"`
	}
	type respondError struct {
		Error string `json:"error"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := s.decode(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse user body error = %v", err)
			s.Respond(w, r, nil, http.StatusBadRequest)
			return
		}

		user, err := s.Store.FindUserByName(req.Username)
		usernameCompare := compareUsername(user.Username, req.Username)
		if usernameCompare == false {
			msg := fmt.Sprint("Username already exist")
			s.Respond(w, r, respondError{
				Error: msg,
			}, http.StatusBadRequest)
			return
		}

		passwordCompare := comparePasswords(req.Password, req.VerifyPassword)
		if passwordCompare == false {
			msg := fmt.Sprint("Password should be similar")
			s.Respond(w, r, respondError{
				Error: msg,
			}, http.StatusBadRequest)
			return
		}
		hashPass, err := utils.HashPassword(req.Password)
		if err != nil {
			log.Printf("Cannot parse user body error = %v", err)
			s.Respond(w, r, nil, http.StatusBadRequest)
			return
		}
		u := &userAccount.User{
			ID:       0,
			Username: req.Username,
			Password: hashPass,
		}

		err = s.Store.CreateUser(u)
		if err != nil {
			log.Printf("Cannot create a user")
			s.Respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		var resp = mapUserToJson(u)
		s.Respond(w, r, resp, http.StatusOK)
	}
}

//mapUserToJson map the Movie in Json
func mapUserToJson(u *userAccount.User) JsonUser {
	return JsonUser{
		ID:       u.ID,
		username: u.Username,
		password: u.Password,
	}
}

//comparePassword compare two given password
func comparePasswords(password string, verifyPassword string) bool {
	if password != verifyPassword {
		return false
	}
	return true
}

//compareUsername compare if username already exist in database username
func compareUsername(usernameDb string, username string) bool {
	if usernameDb == username {
		return false
	}
	return true
}
