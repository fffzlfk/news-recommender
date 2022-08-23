package main

import (
	"news-api/config"
	"news-api/cron"
	"news-api/database"
	"news-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.Init()
	database.Connect()
	cron.Start()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	panic(app.Listen("backend:8080"))
}
