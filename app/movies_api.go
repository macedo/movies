package app

import (
  "net/http"
)

func InitializeMovie() {
  App.Router.HandleFunc("/movies", movieIndex).Methods("GET")
}

func movieIndex(w http.ResponseWriter, r *http.Request) {
  movies, err := App.movie.Get()

  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJson(w, http.StatusOK, movies)
}
