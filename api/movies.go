package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/macedo/movies-api/repository"
)

type MovieAPI struct {
	r repository.MovieRepo
}

func NewMovieAPI(r repository.MovieRepo, root *mux.Router) MovieAPI {
	api := MovieAPI{r: r}

	// register movie API routes
	subrouter := root.PathPrefix("/movies").Subrouter()
	subrouter.HandleFunc("/", api.indexHandler).Methods("GET")

	return api
}

func (api MovieAPI) indexHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := api.r.Get()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, movies)
}
