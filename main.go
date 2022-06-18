package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/handlers"
	"github.com/kuro-vale/kuro-movies-api/middleware"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()
	authMiddleware := middleware.InitJWTMiddleware
	authMiddleware().MiddlewareInit()

	authorized := router.Group("/")
	authorized.Use(authMiddleware().MiddlewareFunc())
	{
		
	}

	// Auth
	router.POST("/auth/signup", handlers.SignUp)
	router.POST("/auth/login", authMiddleware().LoginHandler)
	// Users
	router.GET("/users", handlers.UserIndex)

	router.Run("localhost:8080")
}
