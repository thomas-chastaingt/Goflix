package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/thomas-chastaingt/Goflix/server"
	"github.com/thomas-chastaingt/Goflix/store"
)

//main function to start the server
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

//run permits to start the server in main func
func run() error {
	srv := server.NewServer()
	srv.Store = &store.DbStore{}
	err := srv.Store.Open()
	if err != nil {
		return err
	}

	http.HandleFunc("/", srv.ServHTTP)
	log.Printf("Serving HTTP on port 9000")
	http.ListenAndServe(":9000", nil)
	defer srv.Store.Close()
	return nil
}
