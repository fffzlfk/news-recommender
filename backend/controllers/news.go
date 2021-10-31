package controllers

import (
	"news-api/database"
	"news-api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var MaxNewsNumofPage = myinit()

func myinit() int {
	return viper.GetInt("number.max_news_num_of_page")
}

func GetRecommendNews(c *fiber.Ctx) error {
	MaxNewsNumofPage = viper.GetInt("number.max_news_num_of_page")
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
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

func GetNewsControllerByCategory(category string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")

		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}

		claims := token.Claims.(*jwt.StandardClaims)

		var user models.User

		if err := database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
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

func LikeNews(c *fiber.Ctx) error {
	return nil
}
