package controllers

import (
	"news-api/config"
	"news-api/database"
	"news-api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var maxNewsNumofPage int = config.MaxNewsNumofPage

func NewsRecommendHandler(c *fiber.Ctx) error {
	id := c.Locals("id").(string)

	var user models.User

	res := database.DB.Where("id = ?", id).First(&user)
	if res.Error != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	newsMap := make(map[uint]models.News)

	for i := 0; i < maxNewsNumofPage; i++ {
		category, err := user.GetANewsCategory()
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "internal server error",
			})
		}
		var news models.News
		database.DB.Where("category = ?", category).First(&news)
		newsMap[news.Id] = news
	}

	newsArr := make([]models.News, 0, len(newsMap))

	for _, v := range newsMap {
		newsArr = append(newsArr, v)
	}

	return c.JSON(newsArr)
}

func NewsHandlersByCategory(category string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Locals("id").(string)

		var user models.User

		res := database.DB.Where("id = ?", id).First(&user)
		if res.Error != nil {
			c.Status(fiber.ErrBadRequest.Code)
			return c.JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}

		var newsArr []models.News

		database.DB.Where("category = ?", category).Order("created_at").Find(&newsArr)

		return c.JSON(newsArr)
	}
}

func LikeStateHandler(c *fiber.Ctx) error {
	var user models.User
	userID := c.Locals("id").(string)
	intUserID, _ := strconv.Atoi(userID)
	user.Id = uint(intUserID)

	var news models.News

	newsID := c.Query("news_id")
	intNewsID, _ := strconv.Atoi(newsID)
	news.Id = uint(intNewsID)

	var cnt int
	database.DB.Raw(`SELECT count(*) FROM "users_news" WHERE "news_id"= ?`, news.Id).Scan(&cnt)

	database.DB.Model(&user).Association("LikedNews").Find(&user.LikedNews, []int{int(news.Id)})

	state := false
	if len(user.LikedNews) == 1 {
		state = true
	}

	return c.JSON(fiber.Map{
		"count": cnt,
		"state": state,
	})
}

func LikeNewsHandler(c *fiber.Ctx) error {
	userID := c.Locals("id").(string)
	var user models.User
	res := database.DB.Where("id = ?", userID).First(&user)
	if res.Error != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}

	newsID := c.Query("news_id")
	var news models.News
	res = database.DB.Where("id = ?", newsID).First(&news)
	if res.Error != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}

	action := c.Query("action")

	column := "vis_" + news.Category

	var err error
	if action == "do" {
		database.DB.Model(&user).Update(column, gorm.Expr(column+" + ?", config.Increasement))

		err = database.DB.Model(&user).Association("LikedNews").Append(&news)
	} else if action == "undo" {
		database.DB.Model(&user).Update(column, gorm.Expr(column+" - ?", config.Increasement))

		err = database.DB.Model(&user).Association("LikedNews").Delete(&news)
	} else {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"message": "param action is invalid",
		})
	}

	if err != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
