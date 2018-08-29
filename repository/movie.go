package repository

import (
	"database/sql"

	"github.com/macedo/movies-api/types"
)

type MovieRepository struct {
	db *sql.DB
}

func New(db *sql.DB) MovieRepository {
	return MovieRepository{db: db}
}

func (mr MovieRepository) Get() ([]types.Movie, error) {
	rows, err := mr.db.Query("SELECT * FROM movies")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	movies := []types.Movie{}

	for rows.Next() {
		var m types.Movie
		if err := rows.Scan(&m.ID, &m.Title); err != nil {
			return nil, err
		}

		movies = append(movies, m)
	}

	return movies, nil
}
