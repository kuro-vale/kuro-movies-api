package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null;default:null"`
	Password string `gorm:"not null;default:null"`
}

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       uint      `json:"id"`
	Email    string    `json:"email"`
	JoinDate time.Time `json:"join_date"`
	Links    gin.H     `json:"_links"`
}
