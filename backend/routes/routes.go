package routes

import (
	"news-api/controllers"
	"news-api/models"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.RegisterHandler)

	app.Post("/api/login", controllers.LoginHandler)

	app.Use("/api", controllers.JwtAuthMiddleware())

	app.Get("/api/user", controllers.UserHandler)

	app.Post("/api/logout", controllers.LogoutHandler)

	app.Get("/api/news/recommend", controllers.NewsRecommendHandler)

	for _, category := range models.Categorys {
		app.Get("/api/news/"+category, controllers.NewsHandlersByCategory(category))
	}
}
