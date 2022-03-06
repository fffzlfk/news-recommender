package service

import (
	"log"
	"news-api/config"
	"news-api/database"
	"news-api/models"
	"news-api/utils/similarity"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func updateRecommendList(user models.User, news models.News) (err error) {
	var recentNews []models.News
	err = database.DB.Model(&models.News{}).Find(&recentNews).Order("created_at DESC").Limit(100).Error
	if err != nil {
		return
	}

	err = database.DB.First(&news).Where("news.ID = ?", news.ID).Error
	if err != nil {
		return
	}

	r, closeFunc := similarity.NewRecommender(recentNews)
	defer closeFunc()

	data := make([]models.News, 10)
	copy(data, r.SimOrderNews(news, recentNews))

	for _, v := range data {
		err = database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.RecommendedNews{
			UserID:        user.ID,
			NewsID:        v.ID,
			RecommendedAt: time.Now().Unix(),
		}).Error
		if err != nil {
			return
		}
	}

	return
}

func NewsRecommendService(user *models.User, pageIndex int64) (newsArr []models.News, pageNum int64, err error) {
	pageSize := config.GetPageSize()
	total := database.DB.Model(&user).Association("RecommendedNews").Count()
	pageNum = (total + int64(pageSize) - 1) / int64(pageSize)

	err = database.DB.Raw(
		"SELECT news.* FROM news INNER JOIN recommended_news ON news.id=recommended_news.news_id WHERE recommended_news.user_id=? ORDER BY recommended_news.recommended_at DESC LIMIT ? OFFSET ?",
		user.ID, pageSize, int(pageIndex-1)*pageSize).Find(&newsArr).Error

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

func LikeNewsService(userID, newsID int, action string) (err error) {
	var user models.User
	user.ID = uint(userID)

	var news models.News
	news.ID = uint(newsID)

	if action == "do" {
		err = database.DB.Model(&user).Association("LikedNews").Append(&news)
	} else if action == "undo" {
		err = database.DB.Model(&user).Association("LikedNews").Delete(&news)
	} else {
		err = fiber.NewError(fiber.StatusBadRequest, "param action is invalid")
	}

	go func(user models.User, news models.News) {
		if err := updateRecommendList(user, news); err != nil {
			log.Println(err)
		}
	}(user, news)

	return
}

func ClickNewsService(userID, newsID int) error {
	var user models.User
	user.ID = uint(userID)

	var news models.News
	news.ID = uint(newsID)

	if err := database.DB.Model(&user).Association("ClickedNews").Append(&news); err != nil {
		return err
	}

	go func(user models.User, news models.News) {
		if err := updateRecommendList(user, news); err != nil {
			log.Println(err)
		}
	}(user, news)

	return nil
}

func ColdStartService(userID int, categorys []string) error {
	var user models.User
	user.ID = uint(userID)

	for _, category := range categorys {
		user.ColdStart = append(user.ColdStart, models.CategoryType{
			Category: category,
		})
	}

	if err := database.DB.Model(&user).Association("ColdStart").Replace(user.ColdStart); err != nil {
		return err
	}

	if err := updateRecommendListFromColdStart(user); err != nil {
		return err
	}

	return nil
}

func updateRecommendListFromColdStart(user models.User) error {
	if err := database.DB.First(&user, user.ID).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&user).Association("ColdStart").Find(&user.ColdStart); err != nil {
		return err
	}

	for _, category := range user.ColdStart {
		var newsArr []models.News
		if err := database.DB.Where("category = ?", category.Category).Find(&newsArr).Error; err != nil {
			return err
		}
		for _, news := range newsArr {
			if err := database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.RecommendedNews{
				UserID:        user.ID,
				NewsID:        news.ID,
				RecommendedAt: time.Now().Unix(),
			}).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
