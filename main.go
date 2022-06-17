package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()

	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run("localhost:8080")
}