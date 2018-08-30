package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/macedo/movies-api/types"
)

var fakeMovies = []types.Movie{
	types.Movie{ID: "100", Title: "Rocky III"},
	types.Movie{ID: "200", Title: "Scarface"},
}

type FakeMovieRepository struct{}

func (fake FakeMovieRepository) Get() ([]types.Movie, error) {
	return fakeMovies, nil
}

type ErrorMovieRepository struct{}

func (e ErrorMovieRepository) Get() ([]types.Movie, error) {
	return []types.Movie{}, errors.New("something went wrong")
}

func Test_MovieAPIIndex(t *testing.T) {
	var repo FakeMovieRepository

	api := NewMovieAPI(repo)
	ts := httptest.NewServer(api.Handler())

	defer ts.Close()

	res, err := http.Get(ts.URL + "/movies")

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	var response []types.Movie

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !reflect.DeepEqual(response, fakeMovies) {
		t.Fatalf("expected %v got %v", fakeMovies, response)
	}
}

func Test_MovieAPIIndexWithError(t *testing.T) {
	var repo ErrorMovieRepository

	api := NewMovieAPI(repo)
	ts := httptest.NewServer(api.Handler())

	defer ts.Close()

	res, err := http.Get(ts.URL + "/movies")

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	var response struct{ Error string }

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expectedMessage := "something went wrong"
	if response.Error != expectedMessage {
		t.Fatalf("expected an error message %s got %s", expectedMessage, response.Error)
	}
}
