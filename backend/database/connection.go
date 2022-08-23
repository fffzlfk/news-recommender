package database

import (
	"fmt"
	"log"
	"news-api/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to the database", err)
	}
	DB = conn

	conn.SetupJoinTable(&models.User{}, "RecommendedNews", &models.RecommendedNews{})
	conn.AutoMigrate(&models.User{}, &models.News{}, &models.CategoryType{})
}
