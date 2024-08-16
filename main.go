package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

type userGooseAssociation struct {
	Id      string
	UserId  int
	GooseId int
}
type userGooseDto struct {
	AssociationId string
	Id            int
	Name          string
}

var gooses = []goose{
	{Id: 1, Name: "TheGoldestGoose"},
	{Id: 2, Name: "TheOtherGoose"},
}

var users = []user{
	{Id: 1, Username: "RodrasPT", password: "benfica"},
	{Id: 2, Username: "RocasPT", password: "sporting"},
}

var userGooseAssociations = []userGooseAssociation{}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:userId", getUser)
	router.GET("/gooses", getGooses)
	router.GET("/catch-goose", catchGoose)
	router.GET("/users/:userId/gooses", getUserGeese)
	router.GET("/gooses/:gooseId", getGoose)
	router.DELETE("/users/:userId/gooses/:associationId", releaseGoose)
	router.Run("localhost:8080")

}

// getUsers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	stringId := c.Param("userId")
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

func getGooses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gooses)
}

func catchGoose(c *gin.Context) {
	userId := 1
	randomGooseIndex := rand.Intn(len(gooses))
	chosenGoose := gooses[randomGooseIndex]
	userGoose := userGooseAssociation{uuid.NewString(), userId, chosenGoose.Id}
	userGooseAssociations = append(userGooseAssociations, userGoose)
	c.IndentedJSON(http.StatusOK, chosenGoose)
}

func releaseGoose(c *gin.Context) {
	userId, err := paramToInt(c, "userId")
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

func getUserGeese(c *gin.Context) {
	userId, err := paramToInt(c, "userId")
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	geeseAssociationsForUserId := filter(userGooseAssociations, func(association userGooseAssociation) bool {
		return association.UserId == userId
	})

	var userGeese = transform(geeseAssociationsForUserId, func(association userGooseAssociation) userGooseDto {
		gooseIndex := slices.IndexFunc(gooses, func(goose goose) bool { return association.GooseId == goose.Id })
		userGoose := gooses[gooseIndex]
		var gooseDto = userGooseDto{
			AssociationId: association.Id,
			Id:            userGoose.Id,
			Name:          userGoose.Name,
		}

		return gooseDto
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

func paramToInt(c *gin.Context, urlParam string) (int, error) {
	stringParam := c.Param(urlParam)
	intParam, err := strconv.Atoi(stringParam)

	if err != nil {
		fmt.Printf("parameter '%s' with value '%s' not valid", urlParam, stringParam)
		return 0, err
	}
	return intParam, nil
}
