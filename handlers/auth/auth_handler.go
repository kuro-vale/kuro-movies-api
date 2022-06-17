package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var request models.UserRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": strings.Split(err.Error(), "Error:")[1],
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	newUser := models.User{
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	if err := database.DB.Create(&newUser).Error; err == nil {
		response := UserAssembler(c, newUser)
		c.JSON(http.StatusCreated, response)
		return
	}

	c.Status(http.StatusBadRequest)
}

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
