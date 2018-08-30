package types

// Movie is the struct which maps to movies into database
type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

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

// MovieRepository represents an interface for get movies information
type MovieRepository interface {
	Get() ([]Movie, error)
}
