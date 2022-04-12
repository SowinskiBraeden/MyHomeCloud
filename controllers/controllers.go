package controllers

import (
	"fmt"
	"strings"

	"github.com/SowinskiBraeden/gfbmb/messageBox"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadImage(c *fiber.Ctx) error {
	file, fileError := c.FormFile("image")
	if fileError != nil {
		return c.Status(fiber.StatusInternalServerError).Render("upload", fiber.Map{
			"errorMsg": "Missing File",
		})
	}
	// description := c.FormValue("description")
	// album := c.FormValue("album")

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	err := c.SaveFile(file, fmt.Sprintf("./images/%s", image))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("upload", fiber.Map{
			"errorMsg": messageBox.NewDangerBox("the image could not be saved"),
		})
	}

	// var newImage models.Image
	// newImage.ID = primitive.NewObjectID()
	// newImage.FSName = filename
	// newImage.Description = description
	// newImage.Album =

	return c.Status(fiber.StatusNotImplemented).Render("upload", fiber.Map{
		"errorMsg": messageBox.NewSuccessBox("Success"),
	})
}
