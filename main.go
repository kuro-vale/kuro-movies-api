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
		authorized.PATCH("/actors/:id", handlers.UpdateActor)
		authorized.DELETE("/actors/:id", handlers.DeleteActor)
		// Movies
		authorized.POST("/movies", handlers.StoreMovie)
		authorized.PATCH("/movies/:id", handlers.UpdateMovie)
		authorized.DELETE("/movies/:id", handlers.DeleteMovie)
	}

	// Auth
	router.POST("/auth/signup", handlers.SignUp)
	router.POST("/auth/login", authMiddleware().LoginHandler)
	// Users
	router.GET("/users", handlers.UserIndex)
	router.GET("/users/:id", handlers.ShowUser)
	// Actors
	router.GET("/actors", handlers.ActorIndex)
	router.GET("/actors/:id", handlers.ShowActor)
	router.GET("/actors/:id/movies", handlers.ShowActorMovies)
	// Movies
	router.GET("/movies", handlers.MovieIndex)
	router.GET("/movies/:id", handlers.ShowMovie)
	router.GET("/movies/:id/cast", handlers.ShowMovieCast)

	// GraphQL
	router.GET("/graph", handlers.PlaygroundHandler())
	router.POST("/graph", handlers.GraphqlHandler())

	// Home
	router.LoadHTMLFiles("static/home.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	router.Run(":" + port)
}
