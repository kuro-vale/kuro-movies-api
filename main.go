package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/handlers"
	"github.com/kuro-vale/kuro-movies-api/middleware"
	"github.com/kuro-vale/kuro-movies-api/models"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()
	authMiddleware := middleware.InitJWTMiddleware
	authMiddleware().MiddlewareInit()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	authorized := router.Group("/")
	authorized.Use(authMiddleware().MiddlewareFunc())
	{
		authorized.GET("/jwt/ping", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			user, _ := c.Get("ID")
			c.JSON(200, gin.H{
				"message":   "pong",
				"userClaim": claims["ID"],
				"userID":    user.(models.User).ID,
			})
		})
	}

	router.POST("/auth/signup", handlers.SignUp)
	router.POST("/auth/login", authMiddleware().LoginHandler)

	router.Run("localhost:8080")
}
