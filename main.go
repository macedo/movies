package main

import (
  "fmt"
  "log"
  "os/user"
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
  env := getenv("ENV", "development")

  user, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }

  app := App{
    ENV: map[string]string{
      "DATABASE_HOST": getenv("DATABASE_HOST", "localhost"),
      "DATABASE_NAME": fmt.Sprintf("movies_api_%s", env),
      "DATABASE_PASSWORD": getenv("DATABASE_PASSWORD", ""),
      "DATABASE_URL": getenv("DATABASE_URL", ""),
      "DATABASE_USERNAME": getenv("DATABASE_USERNAME", user.Username),
    },
  }
  app.Initialize()
  app.Run(":8000")
}
