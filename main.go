package main

import (
	"fmt"
	"os"

	"github.com/thomas-chastaingt/Goflix/server"
	"github.com/thomas-chastaingt/Goflix/store"
)

func main() {
	fmt.Println("hello world")
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	srv := server.NewServer()
	srv.Store = &store.DbStore{}
	err := srv.Store.Open()
	if err != nil {
		return err
	}

	movies, err := srv.Store.GetMovies()
	if err != nil {
		return err
	}
	fmt.Printf("movies=%v\n", movies)
	defer srv.Store.Close()
	return nil
}
