package main

import (
	"fmt"
	"math/rand"
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

var userGooseAssociations = []userGoose{}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.GET("/gooses", getGooses)
	router.GET("/catch-goose", catchGoose)
	router.GET("/users/:id/gooses", getUserGooses)
	router.GET("/gooses/:id", getGoose)
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
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var userIndex = slices.IndexFunc(users, func(user user) bool { return user.Id == id })
	c.IndentedJSON(http.StatusOK, users[userIndex])
}

func getGoose(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		fmt.Printf("%s not valid", stringId)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var gooseIndex = slices.IndexFunc(gooses, func(goose goose) bool {return goose.Id == id})
	if  gooseIndex == -1 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	
	c.IndentedJSON(http.StatusOK, gooses[gooseIndex])
}

func getGooses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gooses)
}

func catchGoose(c *gin.Context) {
	userId := 1
	randomGooseIndex := rand.Intn(len(gooses))
	chosenGoose := gooses[randomGooseIndex]
	userGoose := userGoose{userId, chosenGoose.Id}
	userGooseAssociations = append(userGooseAssociations, userGoose)
	c.IndentedJSON(http.StatusOK, chosenGoose)
}

func getUserGooses(c *gin.Context) {
	stringId := c.Param("id")
	userId, err := strconv.Atoi(stringId)

	if err != nil {
		fmt.Printf("%s not valid", stringId)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	geeseAssociationsForUserId := filter(userGooseAssociations, func(gooseLink userGoose) bool {
		return gooseLink.UserId == userId
	})

	var userGeese = transform(geeseAssociationsForUserId, func(link userGoose) goose {
		gooseIndex := slices.IndexFunc(gooses, func(goose goose) bool { return link.GooseId == goose.Id })
		return gooses[gooseIndex]
	})
	c.IndentedJSON(http.StatusOK, userGeese)
}

func transform[T, V any](slice []T, fn func(T) V) []V {
	newSlice := []V{}
	for i := 0; i < len(slice); i++ {
		newSlice = append(newSlice, fn(slice[i]))
	}
	return newSlice
}

func filter[T any](slice []T, funFunction func(T) bool) []T {
	newSlice := []T{}
	for i := 0; i < len(slice); i++ {
		if funFunction(slice[i]) {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}
