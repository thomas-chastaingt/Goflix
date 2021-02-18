package movie

import "fmt"

//Movie define a movie
type Movie struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Duration    int    `db:"duration"`
	TrailerURL  string `db:"trailer_url"`
}

//String display a movie in command line
func (m Movie) String() string {
	return fmt.Sprintf("id=%v, title=%v, release_date=%v, duration=%v, trailer=%v", m.ID, m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
}
