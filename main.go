package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/handlers"
	"github.com/kuro-vale/kuro-movies-api/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	database.ConnectDatabase()
	authMiddleware := middleware.InitJWTMiddleware
	authMiddleware().MiddlewareInit()

	authorized := router.Group("/")
	authorized.Use(authMiddleware().MiddlewareFunc())
	{
		// Users
		authorized.DELETE("/users/:id", handlers.DeleteUser)
		// Actors
		authorized.POST("/actors", handlers.StoreActor)
	}

	// Auth
	router.POST("/auth/signup", handlers.SignUp)
	router.POST("/auth/login", authMiddleware().LoginHandler)
	// Users
	router.GET("/users", handlers.UserIndex)
	router.GET("/users/:id", handlers.ShowUser)
	// Actors
	router.GET("/actors", handlers.ActorIndex)

	router.Run("localhost:" + port)
}
