package database

import (
	"fmt"
	"log"
	"news-api/config"
	"news-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to the database", err)
	}
	DB = conn

	conn.SetupJoinTable(&models.User{}, "LikedNews", &models.LikedNews{})
	conn.AutoMigrate(models.User{})
	conn.AutoMigrate(models.News{})
}
