package service

import (
	"log"
	"news-api/config"
	"news-api/database"
	"news-api/models"
	"news-api/utils/similarity"
	"time"

	"github.com/gofiber/fiber/v2"
)

const maxMotherNews = 10

func updateRecommendList(user models.User) error {
	err := database.DB.Raw(
		"SELECT news.* FROM news INNER JOIN liked_news ON news.id=liked_news.news_id WHERE liked_news.user_id=? ORDER BY liked_news.liked_at DESC LIMIT ?",
		user.ID, maxMotherNews).Find(&user.LikedNews).Error
	if err != nil {
		panic(err)
	}

	if len(user.LikedNews) == 0 {
		return nil
	}

	var recentNews []models.News
	database.DB.Model(&models.News{}).Find(&recentNews).Order("created_at DESC").Limit(100)

	r, closeFunc := similarity.NewRecommender(recentNews)
	defer closeFunc()

	data := make([]models.News, 0)
	for _, likedNews := range user.LikedNews {
		data = append(data, r.SimOrderNews(likedNews, recentNews)...)
	}

	return database.DB.Model(&user).Association("RecommendNews").Replace(data)
}

func NewsRecommendService(user *models.User, pageIndex int64) (newsArr []models.News, pageNum int64, err error) {
	pageSize := config.GetPageSize()
	total := database.DB.Model(&user).Association("RecommendNews").Count()
	pageNum = (total + int64(pageSize) - 1) / int64(pageSize)
	if resErr := database.DB.Offset(int(pageIndex-1) * pageSize).Limit(pageSize).Model(&user).Association("RecommendNews").Find(&newsArr); err != nil {
		err = resErr
	}
	return
}

func NewsByCategoryService(category string, pageIndex int64) (newsArr []models.News, pageNum int64, err error) {
	pageSize := config.GetPageSize()
	var total int64
	res := database.DB.Where("category = ?", category).Model(&models.News{}).Count(&total)
	if res.Error != nil {
		err = res.Error
		return
	}

	pageNum = (total + int64(pageSize) - 1) / int64(pageSize)
	res = database.DB.Offset(int(pageIndex-1)*pageSize).Limit(pageSize).Where("category = ?", category).Order("created_at").Find(&newsArr)
	if res.Error != nil {
		err = res.Error
	}
	return
}

func LikeStateService(userID, newsID int) (state bool, err error) {
	var user models.User
	user.ID = uint(userID)

	var news models.News
	news.ID = uint(newsID)

	err = database.DB.Model(&user).Association("LikedNews").Find(&user.LikedNews, []int{int(news.ID)})
	state = false
	if len(user.LikedNews) == 1 {
		state = true
	}
	return
}

func LikeNewsService(userID, newsID, action string) (err error) {
	var user models.User
	if res := database.DB.Where("id = ?", userID).First(&user); res.Error != nil {
		err = res.Error
		return
	}

	var news models.News
	if res := database.DB.Where("id = ?", newsID).First(&news); res.Error != nil {
		err = res.Error
		return
	}

	if action == "do" {
		database.DB.Create(&models.LikedNews{
			UserID:  user.ID,
			NewsID:  news.ID,
			LikedAt: time.Now().Unix(),
		})
	} else if action == "undo" {
		err = database.DB.Model(&user).Association("LikedNews").Delete(&news)
	} else {
		err = fiber.NewError(fiber.StatusBadRequest, "param action is invalid")
	}

	go func(user models.User) {
		if err := updateRecommendList(user); err != nil {
			log.Println(err)
		}
	}(user)

	return
}
