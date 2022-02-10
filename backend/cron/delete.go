package cron

import (
	"news-api/database"
	"news-api/models"

	"time"
)

// deleteNews delete 3 days ago from table news
func deleteNews() {
	var newsArr []models.News
	database.DB.Where("created_at < ?", time.Now().AddDate(0, 0, -3).Unix()).Find(&newsArr)
	for _, news := range newsArr {
		database.DB.Exec("DELETE FROM liked_news WHERE news_id=?", news.ID)
		database.DB.Exec("DELETE FROM recommend_news WHERE news_id=?", news.ID)
		database.DB.Exec("DELETE FROM clicked_news WHERE news_id=?", news.ID)
		database.DB.Model(&models.News{}).Delete(&news)
	}
}
