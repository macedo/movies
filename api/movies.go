package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/macedo/movies-api/types"
)

// MovieAPI represents the API for movies
type MovieAPI struct {
	r types.MovieRepository
}

// NewMovieAPI creates a MovieAPI
func NewMovieAPI(r types.MovieRepository) MovieAPI {
	return MovieAPI{r: r}
}

// Handler expose service's routes
func (api MovieAPI) Handler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", api.indexHandler).Methods("GET")

	return r
}

func (api MovieAPI) indexHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := api.r.Get()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, movies)
}
