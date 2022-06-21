package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Actor struct {
	gorm.Model
	Name   string  `gorm:"size:50;not null;default:null"`
	Age    uint    `gorm:"check:age >= 18 AND age <= 90;not null;default:null"`
	Gender string  `gorm:"check:gender='Female' OR gender='Male' OR gender='X';not null;default:null"`
	Movies []Movie `gorm:"many2many:cast;"`
}

type StoreActorRequest struct {
	Name   string `json:"name" binding:"required,max=50"`
	Age    uint   `json:"age" binding:"required,lte=90,gte=18"`
	Gender string `json:"gender" binding:"required"`
}

type UpdateActorRequest struct {
	Name   string `json:"name" binding:"max=50"`
	Age    uint   `json:"age" binding:"lte=90,gte=18"`
	Gender string `json:"gender"`
}

type ActorResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Age    uint   `json:"age"`
	Gender string `json:"gender"`
	Links  gin.H  `json:"_links"`
}
