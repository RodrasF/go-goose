package handlers

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func ReleaseGoose(c *gin.Context) {
	userId, err := ParamToInt(c, "userId")
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	associationId := c.Param("associationId")

	gooseAssociationIndex := slices.IndexFunc(userGooseAssociations, func(userGooseAssociation userGooseAssociation) bool {
		return userGooseAssociation.UserId == userId && userGooseAssociation.Id == associationId
	})
	if gooseAssociationIndex == -1 {
		c.Status(http.StatusNoContent)
		return
	}

	userGooseAssociations = slices.Delete(userGooseAssociations, gooseAssociationIndex, gooseAssociationIndex+1)
	c.Status(http.StatusNoContent)
}
