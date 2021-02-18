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
	GetMovieById(id int64) (*movie.Movie, error)
	CreateMovie(m *movie.Movie) error
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

CREATE TABLE IF NOT EXISTS user
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT,
	password TEXT,
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

func (store *DbStore) GetMovieById(id int64) (*movie.Movie, error) {
	var movie = &movie.Movie{}
	err := store.db.Get(movie, "SELECT * FROM movie WHERE id=$1", id)
	if err != nil {
		return movie, nil
	}
	return movie, nil
}

func (Store *DbStore) CreateMovie(m *movie.Movie) error {
	res, err := Store.db.Exec("INSERT INTO movie (title, release_date, duration, trailer_url) VALUES (?,?,?,?)", m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
	if err != nil {
		return err
	}
	m.ID, err = res.LastInsertId()
	return err
}

func (Store *DbStore) FindUser(username string, password string) (bool, error) {
	var count int
	err := Store.db.Get(&count, "SELECT COUNT(id) FROM user WHERE username=$1 AND password=$2", username, password)
	if err != nil {
		return false, nil
	}
	return count == 1, nil
}
