package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/models"
	"github.com/kuro-vale/kuro-movies-api/tools"
)

func MovieIndex(c *gin.Context) {
	pageLimit := 10
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	// Queries
	title := c.Query("title")
	genre := c.Query("genre")
	director := c.Query("director")
	producer := c.Query("producer")

	var count int64
	var movies []models.Movie
	// Query to get the count
	database.DB.Find(&movies, "title LIKE ? AND genre LIKE ? AND director LIKE ? AND producer LIKE ?", "%"+title+"%", "%"+genre+"%", "%"+director+"%", "%"+producer+"%").Count(&count)
	// Query to get the results
	database.DB.Limit(pageLimit).Offset((page-1)*pageLimit).Find(&movies, "title LIKE ? AND genre LIKE ? AND director LIKE ? AND producer LIKE ?", "%"+title+"%", "%"+genre+"%", "%"+director+"%", "%"+producer+"%")

	var response []models.MovieResponse
	for _, movie := range movies {
		movie := movieAssembler(c, movie)
		response = append(response, *movie)
	}

	if len(title) > 0 {
		title = "&title=" + title
	}
	if len(genre) > 0 {
		genre = "&genre=" + genre
	}
	if len(director) > 0 {
		director = "&director=" + director
	}
	if len(producer) > 0 {
		producer = "&producer=" + producer
	}

	links := tools.PaginateIndex(c, page, pageLimit, count, "movies", title, genre, director, producer)

	c.JSON(http.StatusOK, gin.H{
		"data":   response,
		"_links": links,
	})
}

func StoreMovie(c *gin.Context) {
	var request models.StoreMovieRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tools.FormatErr(err.Error()),
		})
		return
	}

	newMovie := models.Movie{
		Title:    request.Title,
		Genre:    request.Genre,
		Price:    fmt.Sprintf("$%v", request.Price),
		Director: request.Director,
		Producer: request.Producer,
	}
	if err := database.DB.Create(&newMovie).Error; err == nil {
		response := movieAssembler(c, newMovie)
		c.JSON(http.StatusCreated, response)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tools.FormatErr(err.Error()),
		})
		return
	}
}

func movieAssembler(c *gin.Context, movie models.Movie) *models.MovieResponse {
	movieResponse := models.MovieResponse{
		ID:       movie.ID,
		Title:    movie.Title,
		Genre:    movie.Genre,
		Price:    movie.Price,
		Director: movie.Director,
		Producer: movie.Producer,
		Links: gin.H{
			"self": gin.H{
				"href": fmt.Sprintf("%s/movies/%d", c.Request.Host, movie.ID),
			},
		},
	}
	return &movieResponse
}
