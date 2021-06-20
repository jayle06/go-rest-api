package main

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/product/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	baseURL := app.Group("api/v1")
	routes.ConfigUserRouter(&baseURL)
	routes.ConfigProductRouter(&baseURL)
	routes.ConfigReviewRouter(&baseURL)
	routes.ConfigCategoryRouter(&baseURL)

	app.Listen(":3000")
}
