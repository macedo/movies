package main

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
	app := App{}
	app.Initialize("go", "gorocks", "localhost", "movies_api_development")
	app.Run(":8000")
}
