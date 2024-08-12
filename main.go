package main

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Goose identifies the structure of a Goose that will be stored in the database table Gooses
type goose struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// User identifies the structure of a User that will be stored in the database table Users
type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	password string
}

// UserGoose identifies the structure of a goose that belongs to a user and will be
// stored in the database table UserGooses
type userGoose struct {
	UserId  int
	GooseId int
}

var gooses = []goose{
	{Id: 1, Name: "TheGoldestGoose"},
	{Id: 2, Name: "TheOtherGoose"},
}

var users = []user{
	{Id: 1, Username: "RodrasPT", password: "benfica"},
	{Id: 2, Username: "RocasPT", password: "sporting"},
}

var usergooses = []userGoose{}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.Run("localhost:8080")
}

// getUsers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		fmt.Printf("%s not valid", stringId)
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}
	
	var userIndex = slices.IndexFunc(users, func(user user) bool { return user.Id == id })
	c.IndentedJSON(http.StatusOK, users[userIndex])
}
