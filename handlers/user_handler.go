package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/models"
)

func UserAssembler(c *gin.Context, user models.User) *models.UserResponse {
	userResponse := models.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		JoinDate: user.CreatedAt,
		Links: gin.H{
			"self": gin.H{
				"href": fmt.Sprintf("%s/users/%d", c.Request.Host, user.ID),
			},
		},
	}
	return &userResponse
}
