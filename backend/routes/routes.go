package routes

import (
	"news-api/controllers"
	"news-api/models"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)

	app.Post("/api/login", controllers.Login)

	app.Get("/api/user", controllers.User)

	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/news/recommend", controllers.GetRecommendNews)

	for _, category := range models.Categorys {
		app.Get("/api/news/"+category, controllers.GetNewsControllerByCategory(category))
	}
}
