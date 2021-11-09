package controllers

import (
	"log"
	"news-api/config"
	"news-api/database"
	"news-api/models"
	"news-api/utils/similarity"
	"strconv"

	"github.com/gofiber/fiber/v2"
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

	var data []models.News
	if err := database.DB.Model(&user).Association("RecommendNews").Find(&data); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(data)
}

func updateRecommendList(user models.User) error {
	database.DB.Model(&user).Association("LikedNews").Find(&user.LikedNews)
	if len(user.LikedNews) == 0 {
		todo()
	}

	var allNews []models.News
	database.DB.Model(&models.News{}).Find(&allNews).Order("created_at DESC").Limit(100)

	data := make([]models.News, 0)

	for _, likedNews := range user.LikedNews {
		r := similarity.NewRecommend(*likedNews)
		data = append(data, r.SimOrderNews(allNews)...)
	}

	return database.DB.Model(&user).Association("RecommendNews").Replace(data)
}

func todo() {}

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

	database.DB.Model(&user).Association("LikedNews").Find(&user.LikedNews, []int{int(news.Id)})

	state := false
	if len(user.LikedNews) == 1 {
		state = true
	}

	return c.JSON(fiber.Map{
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

	var err error
	if action == "do" {
		err = database.DB.Model(&user).Association("LikedNews").Append(&news)
	} else if action == "undo" {
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

	go func(user models.User) {
		if err := updateRecommendList(user); err != nil {
			log.Println(err)
		}
	}(user)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
