package database

import (
	"os"

	"github.com/kuro-vale/kuro-movies-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Actor{})
	database.AutoMigrate(&models.Movie{})

	DB = *database
}
