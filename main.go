package main

import (
	"go-goose/api"
	"go-goose/data/db"
	"go-goose/data/repository"
)

func main() {
	DB := db.Connect()

	repository := repository.New(DB)
	api.HandleRequests(repository)

	db.CloseConnection(DB)
}
