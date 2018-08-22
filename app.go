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
}

func (a *App) Initialize(user, password, host, dbname string) {
  //connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)
	connStr := fmt.Sprintf("user=macedo dbname=movies_api_development sslmode=disable")

	var err error

	a.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
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
