package models

// Goose identifies the structure of a Goose that will be stored in the database table Gooses
type Goose struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
