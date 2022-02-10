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

	newsAPI := app.Group("/api/news")
	newsAPI.Get("/recommend", controllers.NewsRecommendHandler)
	for _, category := range models.Categorys {
		newsAPI.Get("/"+category, controllers.NewsHandlersByCategory(category))
	}

	likeAPI := app.Group("/api/like")
	likeAPI.Get("/action", controllers.LikeNewsHandler)
	likeAPI.Get("/get", controllers.LikeStateHandler)

	app.Get("/api/click", controllers.ClickNewsHandler)
}
