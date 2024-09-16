package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetUser(c *gin.Context) {
	stringId := c.Param("userId")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		fmt.Printf("%s not valid", stringId)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, h.Repository.GetUserById(id))
}
