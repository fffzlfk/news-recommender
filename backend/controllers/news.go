package controllers

import (
	"news-api/database"
	"news-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var MaxNewsNumofPage = myinit()

func myinit() int {
	return viper.GetInt("number.max_news_num_of_page")
}

func NewsRecommendHandler(c *fiber.Ctx) error {
	MaxNewsNumofPage = viper.GetInt("number.max_news_num_of_page")
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

	for i := 0; i < MaxNewsNumofPage; i++ {
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

func LikeNewsHandler(c *fiber.Ctx) error {
	return nil
}
