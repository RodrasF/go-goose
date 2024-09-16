package handlers

import (
	"go-goose/data/repository"
)

type Handler struct {
	Repository *repository.Repository
}

func New(repo *repository.Repository) Handler {
	return Handler{repo}
}
