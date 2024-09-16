package handlers

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGoose(c *gin.Context) {
	stringId := c.Param("gooseId")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		fmt.Printf("%s not valid", stringId)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var gooseIndex = slices.IndexFunc(gooses, func(goose goose) bool { return goose.Id == id })
	if gooseIndex == -1 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, gooses[gooseIndex])
}
