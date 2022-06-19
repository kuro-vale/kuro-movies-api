package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/models"
	"github.com/kuro-vale/kuro-movies-api/tools"
	"golang.org/x/crypto/bcrypt"
)

func UserIndex(c *gin.Context) {
	pageLimit := 5
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	email := c.Query("email")
	var count int64
	var users []models.User
	// Query to get the count
	database.DB.Find(&users, "email LIKE ?", "%"+email+"%").Count(&count)
	// Query to get the results
	database.DB.Limit(pageLimit).Offset((page-1)*pageLimit).Find(&users, "email LIKE ?", "%"+email+"%")

	var response []models.UserResponse
	for _, user := range users {
		user := userAssembler(c, user)
		response = append(response, *user)
	}

	if len(email) > 0 {
		email = "&email=" + email
	}

	links := tools.PaginateIndex(c, page, pageLimit, count, "users", email)

	c.JSON(http.StatusOK, gin.H{
		"data":   response,
		"_links": links,
	})
}

func SignUp(c *gin.Context) {
	var request models.UserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tools.FormatErr(err.Error()),
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	newUser := models.User{
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	if err := database.DB.Create(&newUser).Error; err == nil {
		response := userAssembler(c, newUser)
		c.JSON(http.StatusCreated, response)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tools.FormatErr(err.Error()),
		})
		return
	}
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err == nil {
		response := userAssembler(c, user)
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("ID")

	var userToDelete models.User
	if err := database.DB.First(&userToDelete, "id = ?", id).Error; err == nil {
		if userToDelete.ID == user.(models.User).ID {
			database.DB.Delete(&userToDelete)
			c.Status(http.StatusNoContent)
			return
		} else {
			c.Status(http.StatusForbidden)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
	})
}

func userAssembler(c *gin.Context, user models.User) *models.UserResponse {
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
