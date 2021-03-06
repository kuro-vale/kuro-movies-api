package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/models"
	"github.com/kuro-vale/kuro-movies-api/tools"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	var cast []models.Actor
	for _, actorRequest := range request.Cast {
		newActor := models.Actor{
			Name:   actorRequest.Name,
			Age:    actorRequest.Age,
			Gender: cases.Title(language.Und).String(actorRequest.Gender),
		}
		if err := database.DB.FirstOrCreate(&newActor, "name = ?", newActor.Name).Error; err == nil {
			cast = append(cast, newActor)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": tools.FormatErr(err.Error()),
			})
			return
		}
	}

	newMovie := models.Movie{
		Title:    request.Title,
		Genre:    request.Genre,
		Price:    fmt.Sprintf("$%v", request.Price),
		Director: request.Director,
		Producer: request.Producer,
		Actors:   cast,
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

func ShowMovie(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie
	if err := database.DB.First(&movie, "id = ?", id).Error; err == nil {
		response := movieAssembler(c, movie)
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "movie not found",
	})
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var request models.UpdateMovieRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tools.FormatErr(err.Error()),
		})
		return
	}

	var cast []models.Actor
	for _, actorRequest := range request.Cast {
		newActor := models.Actor{
			Name:   actorRequest.Name,
			Age:    actorRequest.Age,
			Gender: cases.Title(language.Und).String(actorRequest.Gender),
		}
		if err := database.DB.FirstOrCreate(&newActor, "name = ?", newActor.Name).Error; err == nil {
			cast = append(cast, newActor)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": tools.FormatErr(err.Error()),
			})
			return
		}
	}

	updatedMovie := models.Movie{
		Title:    request.Title,
		Genre:    request.Genre,
		Price:    fmt.Sprintf("$%v", request.Price),
		Director: request.Director,
		Producer: request.Producer,
	}

	var movie models.Movie
	if err := database.DB.First(&movie, "id = ?", id).Error; err == nil {
		if request.Price == 0 {
			updatedMovie.Price = movie.Price
		}
		if err := database.DB.Model(&movie).Updates(updatedMovie).Error; err == nil {
			database.DB.Model(&movie).Association("Actors").Append(cast)
			response := movieAssembler(c, movie)
			c.JSON(http.StatusOK, response)
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": tools.FormatErr(err.Error()),
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "movie not found",
	})
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	var movieToDelete models.Movie
	if err := database.DB.First(&movieToDelete, "id = ?", id).Error; err == nil {
		database.DB.Delete(&movieToDelete)
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "movie not found",
	})
}

func ShowMovieCast(c *gin.Context) {
	id := c.Param("id")

	pageLimit := 10
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	name := c.Query("name")
	var count int64
	var cast []models.Actor
	// Query to get the count
	database.DB.Joins("JOIN public.\"cast\" AS c ON c.actor_id = actors.id AND c.movie_id = ?", id).Find(&cast, "name LIKE ?", "%"+name+"%").Count(&count)
	// Query to get the results
	database.DB.Limit(pageLimit).Offset((page-1)*pageLimit).Joins("JOIN public.\"cast\" AS c ON c.actor_id = actors.id AND c.movie_id = ?", id).Find(&cast, "name LIKE ?", "%"+name+"%")

	var response []models.ActorResponse
	for _, actor := range cast {
		actor := actorAssembler(c, actor)
		response = append(response, *actor)
	}

	if len(name) > 0 {
		name = "&name=" + name
	}

	links := tools.PaginateIndex(c, page, pageLimit, count, "movies/"+id+"/cast", name)

	c.JSON(http.StatusOK, gin.H{
		"data":   response,
		"_links": links,
	})
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
			"cast": gin.H{
				"href": fmt.Sprintf("%s/movies/%d/cast", c.Request.Host, movie.ID),
			},
		},
	}
	return &movieResponse
}
