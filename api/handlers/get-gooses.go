package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGooses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gooses)
}
