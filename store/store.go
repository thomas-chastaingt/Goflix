package store

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/thomas-chastaingt/Goflix/movie"
)

type Store interface {
	Open() error
	Close() error

	GetMovies() ([]*movie.Movie, error)
}

type DbStore struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS movie
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	release_date TEXT,
	duration INTEGER,
	trailer_url TEXT
);
`

func (store *DbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "goflix.db")
	if err != nil {
		return err
	}
	log.Println("Connected to Db")
	db.MustExec(schema)
	store.db = db
	return nil
}

func (store *DbStore) Close() error {
	return store.db.Close()
}

func (store *DbStore) GetMovies() ([]*movie.Movie, error) {
	var movies []*movie.Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return movies, err
	}
	return movies, nil
}
