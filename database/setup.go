package database

import (
	"github.com/kuro-vale/kuro-movies-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Actor{})
	database.AutoMigrate(&models.Movie{})

	DB = *database
}
