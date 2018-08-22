package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
  ENV    map[string] string
}

func (a *App) Initialize() {
	a.initializeDB()
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeDB() {
  var connStr string
  var err error

  if a.ENV["DATABASE_URL"] != "" {
    connStr = a.ENV["DATABASE_URL"]
  } else {
    connStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", a.ENV["DATABASE_USERNAME"], a.ENV["DATABASE_PASSWORD"], a.ENV["DATABASE_HOST"],  a.ENV["DATABASE_NAME"])
  }

  fmt.Println(connStr)

  a.DB, err = sql.Open("postgres", connStr)
	if err != nil {
	  	log.Fatal(err)
	}
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/movies", a.getMovies).Methods("GET")
	a.Router.HandleFunc("/movies", a.createMovie).Methods("POST")
}

func (a *App) getMovies(w http.ResponseWriter, r *http.Request) {
  var m movie
  movies, err := m.getMovies(a.DB)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  
  respondWithJson(w, http.StatusOK, movies)
}

func (a *App) createMovie(w http.ResponseWriter, r *http.Request) {
	var m movie
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := m.createMovie(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, m)
}
