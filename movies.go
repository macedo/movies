package main

import (
	"database/sql"
	"fmt"
)

type movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (m *movie) getMovies(db *sql.DB) ([]movie, error) {
  statement := fmt.Sprintf("SELECT * FROM movies")
  rows, err := db.Query(statement)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  movies := []movie{}

  for rows.Next() {
    var m movie
    if err := rows.Scan(&m.ID, &m.Title); err != nil {
      return nil, err
    }

    movies = append(movies, m)
  }
  return movies, nil
}

func (m *movie) createMovie(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO movies(title) VALUES('%s') RETURNING id", m.Title)
	err := db.QueryRow(statement).Scan(&m.ID)

	if err != nil {
		return err
	}

	//err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&m.ID)

	//if err != nil {
		//return err
	//}

	return nil
}
