package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/handlers"
	"github.com/kuro-vale/kuro-movies-api/middleware"
)

func main() {
	port:= os.Getenv("PORT")
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
	}

	// Auth
	router.POST("/auth/signup", handlers.SignUp)
	router.POST("/auth/login", authMiddleware().LoginHandler)
	// Users
	router.GET("/users", handlers.UserIndex)
	router.GET("/users/:id", handlers.ShowUser)

	router.Run("localhost:"+port)
}
