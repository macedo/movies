package repository

import (
	"database/sql"

	"github.com/macedo/movies-api/types"
)

//MovieRepo comunicates with database in order to manage movies entries
type MovieRepo struct {
	db *sql.DB
}

// New create a MovieRepository
func New(db *sql.DB) MovieRepo {
	return MovieRepo{db: db}
}

// Get retuns a list of movies
func (mr MovieRepo) Get() ([]types.Movie, error) {
	rows, err := mr.db.Query("SELECT * FROM movies")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var movies []types.Movie

	for rows.Next() {
		var m types.Movie
		if err := rows.Scan(&m.ID, &m.Title); err != nil {
			return nil, err
		}

		movies = append(movies, m)
	}

	return movies, nil
}
