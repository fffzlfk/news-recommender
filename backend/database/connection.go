package database

import (
	"fmt"
	"log"
	"news-api/models"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		viper.Get("pgsql.host"),
		viper.Get("pgsql.user"),
		viper.Get("pgsql.password"),
		viper.Get("pgsql.dbname"),
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to the database", err)
	}
	DB = conn

	conn.AutoMigrate(models.User{})
	conn.AutoMigrate(models.News{})

	// delete old news
	deleteNews()

	// add new news
	addNews()
}

// deleteNews delete 3 days ago from table news
func deleteNews() {
	DB.Where("created_at < ?", time.Now().AddDate(0, 0, -3).Unix()).Delete(models.News{})
}

func addNews() {
	MaxNewsNumofDB := viper.GetInt64("number.max_news_num_of_db")
	var cnt int64
	DB.Model(models.News{}).Count(&cnt)
	if cnt < MaxNewsNumofDB {
		crawlNew()
	}
}
