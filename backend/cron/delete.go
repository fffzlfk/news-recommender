package cron

import (
	"news-api/database"
	"news-api/models"

	"time"
)

// deleteNews delete 3 days ago from table news
func deleteNews() {
	database.DB.Where("created_at < ?", time.Now().AddDate(0, 0, -3).Unix()).Delete(models.News{})
}
