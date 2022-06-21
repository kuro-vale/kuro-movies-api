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

func ActorIndex(c *gin.Context) {
	pageLimit := 10
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	name := c.Query("name")
	var count int64
	var actors []models.Actor
	// Query to get the count
	database.DB.Find(&actors, "name LIKE ?", "%"+name+"%").Count(&count)
	// Query to get the results
	database.DB.Limit(pageLimit).Offset((page-1)*pageLimit).Find(&actors, "name LIKE ?", "%"+name+"%")

	var response []models.ActorResponse
	for _, actor := range actors {
		actor := actorAssembler(c, actor)
		response = append(response, *actor)
	}

	if len(name) > 0 {
		name = "&name=" + name
	}

	links := tools.PaginateIndex(c, page, pageLimit, count, "actors", name)

	c.JSON(http.StatusOK, gin.H{
		"data":   response,
		"_links": links,
	})
}

func StoreActor(c *gin.Context) {
	var request models.StoreActorRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tools.FormatErr(err.Error()),
		})
		return
	}

	newActor := models.Actor{
		Name:   request.Name,
		Age:    request.Age,
		Gender: cases.Title(language.Und).String(request.Gender),
	}
	if err := database.DB.Create(&newActor).Error; err == nil {
		response := actorAssembler(c, newActor)
		c.JSON(http.StatusCreated, response)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tools.FormatErr(err.Error()),
		})
		return
	}
}

func ShowActor(c *gin.Context) {
	id := c.Param("id")

	var actor models.Actor
	if err := database.DB.First(&actor, "id = ?", id).Error; err == nil {
		response := actorAssembler(c, actor)
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "actor not found",
	})
}

func UpdateActor(c *gin.Context) {
	id := c.Param("id")
	var request models.UpdateActorRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tools.FormatErr(err.Error()),
		})
		return
	}
	updatedActor := models.Actor{
		Name:   request.Name,
		Age:    request.Age,
		Gender: cases.Title(language.Und).String(request.Gender),
	}

	var actor models.Actor
	if err := database.DB.First(&actor, "id = ?", id).Error; err == nil {
		if err := database.DB.Model(&actor).Updates(updatedActor).Error; err == nil {
			response := actorAssembler(c, actor)
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
		"message": "actor not found",
	})
}

func DeleteActor(c *gin.Context) {
	id := c.Param("id")

	var actorToDelete models.Actor
	if err := database.DB.First(&actorToDelete, "id = ?", id).Error; err == nil {
		database.DB.Delete(&actorToDelete)
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "actor not found",
	})
}

func ShowActorMovies(c *gin.Context) {
	id := c.Param("id")

	var actor models.Actor
	if err := database.DB.Preload("Movies").First(&actor, "id = ?", id).Error; err == nil {
		var response []models.MovieResponse
		for _, movie := range actor.Movies {
			movie := movieAssembler(c, movie)
			response = append(response, *movie)
		}
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "actor not found",
	})
}

func actorAssembler(c *gin.Context, actor models.Actor) *models.ActorResponse {
	actorResponse := models.ActorResponse{
		ID:     actor.ID,
		Name:   actor.Name,
		Age:    actor.Age,
		Gender: actor.Gender,
		Links: gin.H{
			"self": gin.H{
				"href": fmt.Sprintf("%s/actors/%d", c.Request.Host, actor.ID),
			},
		},
	}
	return &actorResponse
}
