package api

import (
	"database/sql"
	"net/http"

	l4g "github.com/alecthomas/log4go"
	"github.com/gorilla/mux"
	"github.com/macedo/movies-api/repository"
)

type Application struct {
	r        *mux.Router
	movieAPI MovieAPI
}

func New(db *sql.DB) *Application {
	l4g.Info("App Initializing...")

	app := &Application{
		r: mux.NewRouter().StrictSlash(true),
	}

	repo := repository.New(db)
	app.movieAPI = NewMovieAPI(repo, app.r)

	return app
}

func (a *Application) Handler() http.Handler {
	return a.r
}
