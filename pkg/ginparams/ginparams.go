package ginparams

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IntParam(c *gin.Context, paramKey string) (int, error) {
	stringParam := c.Param(paramKey)
	intParam, err := strconv.Atoi(stringParam)

	if err != nil {
		fmt.Printf("Parameter '%s' with value '%s' is not a valid int.", paramKey, stringParam)
		return 0, err
	}
	return intParam, nil
}
