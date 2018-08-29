package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/user"

	_ "github.com/lib/pq"
	"github.com/macedo/movies-api/api"

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

type Config struct {
	Database Database
	Env      string
}

type Database struct {
	Username string
	Password string
	Host     string
	URL      string
}

func main() {
	var c Config
	_ = envconfig.Process("", &c)

	var connStr string

	if c.Database.URL != "" {
		connStr = c.Database.URL
	} else {
		var env string
		if env = c.Env; env == "" {
			env = "development"
		}
		dbname := fmt.Sprintf("movies_api_%s", env)

		var dbusername string
		if c.Database.Username == "" {
			user, err := user.Current()
			if err != nil {
				log.Fatal(err)
			}
			dbusername = user.Username
		} else {
			dbusername = c.Database.Username
		}
		connStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbusername, c.Database.Password, c.Database.Host, dbname)
	}

	fmt.Println(connStr)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8000", api.New(dbConn).Handler()); err != nil {
		log.Fatal(err)
	}
}
