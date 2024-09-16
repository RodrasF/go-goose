package handlers

import (
	"go-goose/models"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) CatchGoose(c *gin.Context) {
	//TODO: Remove hard-coded user id
	userId := 1
	userGooseAssociations := h.Repository.GetUserGooseAssociations(userId)
	gooses := h.Repository.GetGooses()

	randomGooseIndex := rand.Intn(len(gooses))
	chosenGoose := gooses[randomGooseIndex]
	userGoose := models.UserGooseAssociation{
		Id:      uuid.NewString(),
		UserId:  userId,
		GooseId: chosenGoose.Id,
	}

	// TODO: Does this actually change the userGooseAssociations in the database?
	userGooseAssociations = append(userGooseAssociations, userGoose)

	c.IndentedJSON(http.StatusOK, chosenGoose)
}
