package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/macedo/movies-api/repository"
)

// MovieAPI represents the API for movies
type MovieAPI struct {
	r repository.MovieRepository
}

// NewMovieAPI creates a MovieAPI and register movie endpoints on root router
func NewMovieAPI(r repository.MovieRepository, root *mux.Router) MovieAPI {
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

	respondWithJSON(w, http.StatusOK, movies)
}
