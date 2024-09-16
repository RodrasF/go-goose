package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
