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
