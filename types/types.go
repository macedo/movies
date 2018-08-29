package types

// Movie is the struct which maps to movies into database
type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
