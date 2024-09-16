package api

import (
	"go-goose/api/handlers"
	"go-goose/data/repository"

	"github.com/gin-gonic/gin"
)

func HandleRequests(repo *repository.Repository) {
	h := handlers.New(repo)
	router := gin.Default()

	addRoutes(router, h)

	router.Run("localhost:8080")
}

func addRoutes(router *gin.Engine, h handlers.Handler) *gin.Engine {
	router.GET("/users", h.GetUsers)
	router.GET("/users/:userId", h.GetUser)
	router.GET("/gooses", h.GetGooses)
	router.GET("/catch-goose", h.CatchGoose)
	router.GET("/users/:userId/gooses", h.GetUserGeese)
	router.GET("/gooses/:gooseId", h.GetGoose)
	router.DELETE("/users/:userId/gooses/:associationId", h.ReleaseGoose)

	return router
}
