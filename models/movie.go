package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title    string `gorm:"size:150;not null;default:null"`
	Genre    string `gorm:"size:100;not null;default:null"`
	Price    string `gorm:"size:10;not null;default:null"`
	Director string `gorm:"size:50;not null;default:null"`
	Producer string `gorm:"size:50;not null;default:null"`
}

type StoreMovieRequest struct {
	Title    string `json:"title" binding:"required, max=150"`
	Genre    string `json:"genre" binding:"required, max=100"`
	Price    string `json:"price" binding:"required, max=10"`
	Director string `json:"director" binding:"required, max=50"`
	Producer string `json:"producer" binding:"required, max=50"`
}

type UpdateMovieRequest struct {
	Title    string `json:"title" binding:"max=150"`
	Genre    string `json:"genre" binding:"max=100"`
	Price    string `json:"price" binding:"max=10"`
	Director string `json:"director" binding:"max=50"`
	Producer string `json:"producer" binding:"max=50"`
}

type MovieResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Price    string `json:"price"`
	Director string `json:"director"`
	Producer string `json:"producer"`
	Links    gin.H  `json:"_links"`
}
