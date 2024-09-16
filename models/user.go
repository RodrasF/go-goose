package models

// User identifies the structure of a User that will be stored in the database table Users
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
