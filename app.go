package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/macedo/movies-api/repository"
)

type App struct {
	DB  *sql.DB
	ENV map[string]string

	movie repository.MovieRepo
}

func (a *App) Initialize() {
	a.initializeDB()

	a.movie = repository.New(a.DB)
}

func (a *App) initializeDB() {
	var connStr string
	var err error

	if a.ENV["DATABASE_URL"] != "" {
		connStr = a.ENV["DATABASE_URL"]
	} else {
		connStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", a.ENV["DATABASE_USERNAME"], a.ENV["DATABASE_PASSWORD"], a.ENV["DATABASE_HOST"], a.ENV["DATABASE_NAME"])
	}

	fmt.Println(connStr)

	a.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Handler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/movies", a.getMovies).Methods("GET")
	//r.HandleFunc("/movies", a.createMovie).Methods("POST")

	return r
}

func (a *App) getMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := a.movie.Get()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, movies)
}

//func (a *App) createMovie(w http.ResponseWriter, r *http.Request) {
	//var m repository.MovieRepo
	//decoder := json.NewDecoder(r.Body)
	//if err := decoder.Decode(&m); err != nil {
		//respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		//return
	//}
	//defer r.Body.Close()

	//if err := m.createMovie(a.DB); err != nil {
		//respondWithError(w, http.StatusInternalServerError, err.Error())
		//return
	//}

	//respondWithJson(w, http.StatusCreated, m)
//}
