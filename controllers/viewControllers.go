package controllers

import (
	"github.com/SowinskiBraeden/gfbmb/messageBox"
	"github.com/gofiber/fiber/v2"
)

// Render Index
func MainPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func UploadImageView(c *fiber.Ctx) error {
	return c.Render("upload", fiber.Map{
		"errorMsg": messageBox.EmptyMessageBox(),
	})
}
