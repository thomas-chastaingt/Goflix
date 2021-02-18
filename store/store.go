package store

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/thomas-chastaingt/Goflix/movie"
	userAccount "github.com/thomas-chastaingt/Goflix/user"
)

//Store implement store methods
type Store interface {
	Open() error
	Close() error

	GetMovies() ([]*movie.Movie, error)
	GetMovieById(id int64) (*movie.Movie, error)
	CreateMovie(m *movie.Movie) error

	CreateUser(u *userAccount.User) error
	FindUserByName(username string) (*userAccount.User, error)
}

//DbStore define the database
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
	password TEXT
);
`

//Open permits to open database
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

//Close permits to close database
func (store *DbStore) Close() error {
	return store.db.Close()
}

//GetMovies permits to get all movies from database
//RETURN all movies
func (store *DbStore) GetMovies() ([]*movie.Movie, error) {
	var movies []*movie.Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return movies, err
	}
	return movies, nil
}

//GetMovieById permits to get specific movie from database
//RETURN a movie
func (store *DbStore) GetMovieById(id int64) (*movie.Movie, error) {
	var movie = &movie.Movie{}
	err := store.db.Get(movie, "SELECT * FROM movie WHERE id=$1", id)
	if err != nil {
		return movie, nil
	}
	return movie, nil
}

//CreateMovie permits to create a new movie from database
func (store *DbStore) CreateMovie(m *movie.Movie) error {
	res, err := store.db.Exec("INSERT INTO movie (title, release_date, duration, trailer_url) VALUES (?,?,?,?)", m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
	if err != nil {
		return err
	}
	m.ID, err = res.LastInsertId()
	return err
}

/*************************************** User methods ***************************************/

//FindUserByName permits to get a specific user in databse
func (store *DbStore) FindUserByName(username string) (*userAccount.User, error) {
	var user = &userAccount.User{}
	err := store.db.Get(user, "SELECT * FROM user WHERE username=$1", username)
	if err != nil {
		return user, nil
	}
	return user, nil
}

//CreateUser permits to create a new user in database
func (store *DbStore) CreateUser(u *userAccount.User) error {
	res, err := store.db.Exec("INSERT INTO user (username, password) VALUES (?,?)", u.Username, u.Password)
	if err != nil {
		return err
	}
	u.ID, err = res.LastInsertId()
	return nil
}
