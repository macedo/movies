package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/macedo/movies-api/api"
	"github.com/macedo/movies-api/types"

	"github.com/kelseyhightower/envconfig"
)

//func GetMovieInfo() {
//response, err := http.Get("https://api.themoviedb.org/3/search/movie?api_key=6eadd23fe9bae68f2db394ebf0e6bb48&language=pt-BT&query=senhor%20dos%20aneis&page=1&include_adult=false")

//if err != nil {
//fmt.Print(err.Error())
//os.Exit(1)
//}

//responseData, err := ioutil.ReadAll(response.Body)
//if err != nil {
//log.Fatal(err)
//}

//fmt.Println(string(responseData))
//}

func main() {
	var c types.Config

	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err)
	}

	conn, err := NewConnection(c)
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := sql.Open("postgres", conn.PostgresURI())
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8000", api.New(dbConn).Handler()); err != nil {
		log.Fatal(err)
	}
}
