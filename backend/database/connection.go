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
	c := config.GetDatabaseConfigurations()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		c.Host,
		c.User,
		c.Password,
		c.DBName,
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to the database", err)
	}
	DB = conn

	conn.SetupJoinTable(&models.User{}, "RecommendedNews", &models.RecommendedNews{})
	conn.AutoMigrate(models.User{}, models.News{})
}
