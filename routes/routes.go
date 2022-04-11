package routes

import (
	"github.com/SowinskiBraeden/MyHomeCloud/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Render Handlers
	app.Get("/", controllers.MainPage)
	app.Get("/uploadImage", controllers.UploadImageView)

	// API Handlers
	app.Post("/uploadImage", controllers.UploadImage)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
