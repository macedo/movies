package app

import (
  "database/sql"
  "net/http"
  l4g "github.com/alecthomas/log4go"
  "github.com/gorilla/mux"
  "github.com/macedo/movies-api/repository"
)

type Application struct {
  Router *mux.Router
  DB  *sql.DB

  movie repository.MovieRepo
}

var App *Application

func New(db *sql.DB) Application {
  l4g.Info("App Initializing...")

  App = &Application{
    DB: db,
    Router: NewRouter(),
  }

  App.initializeRepos()
  App.initializeRoutes()

  return *App
}

func (a *Application) Handler() http.Handler {
  return a.Router
}

func (a *Application) initializeRepos() {
  App.movie = repository.New(App.DB)
}

func (a *Application) initializeRoutes() {
  InitializeMovie()
}
