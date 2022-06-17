package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/handlers/auth"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/auth/signup", auth.SignUp)

	router.Run("localhost:8080")
}
