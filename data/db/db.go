package db

import (
	"fmt"
	"go-goose/models"
)

type DB struct {
	Users                 []models.User
	Gooses                []models.Goose
	UserGooseAssociations []models.UserGooseAssociation
}

var gooses = []models.Goose{
	{Id: 1, Name: "TheGoldestGoose"},
	{Id: 2, Name: "TheOtherGoose"},
}

var users = []models.User{
	{Id: 1, Username: "RodrasPT"},
	{Id: 2, Username: "RocasPT"},
}

var userGooseAssociations = []models.UserGooseAssociation{}

func Connect() *DB {
	db := &DB{
		Users:                 users,
		Gooses:                gooses,
		UserGooseAssociations: userGooseAssociations,
	}
	return db
}

func CloseConnection(db *DB) {
	fmt.Println("Closed db connection!")
}
