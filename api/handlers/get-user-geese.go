package handlers

import (
	"go-goose/models"
	"go-goose/pkg/ginparams"
	"go-goose/pkg/sliceutils"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetUserGeese(c *gin.Context) {
	userId, err := ginparams.IntParam(c, "userId")
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	geeseAssociationsForUserId := repo.GetUserGooseAssociations(userId)

	var userGeese = sliceutils.Transform(geeseAssociationsForUserId, func(association models.UserGooseAssociation) models.UserGooseDto {
		gooseIndex := slices.IndexFunc(repo.GetGooses(), func(goose models.Goose) bool { return association.GooseId == goose.Id })
		userGoose := gooses[gooseIndex]
		var gooseDto = models.UserGooseDto{
			AssociationId: association.Id,
			Id:            userGoose.Id,
			Name:          userGoose.Name,
		}

		return gooseDto
	})

	c.IndentedJSON(http.StatusOK, userGeese)
}
