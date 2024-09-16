package repository

import (
	"go-goose/data/db"
	"go-goose/models"
)

type Repository struct {
	db *db.DB
}

func New(db *db.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) GetUsers() []models.User {
	return repo.db.Users
}

func (repo *Repository) GetUserById(id int) *models.User {
	for _, user := range repo.db.Users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func (repo *Repository) GetGooses() []models.Goose {
	return repo.db.Gooses
}

func (repo *Repository) GetGooseById(id int) *models.Goose {
	for _, goose := range repo.db.Gooses {
		if goose.Id == id {
			return &goose
		}
	}
	return nil
}

func (repo *Repository) GetUserGooseAssociations(userId int) []models.UserGooseAssociation {
	var userGooses []models.UserGooseAssociation
	for _, userGoose := range repo.db.UserGooseAssociations {
		if userGoose.UserId == userId {
			userGooses = append(userGooses, userGoose)
		}
	}
	return userGooses
}
