package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/poonyawat0511/go-fiber/services"
)

func RegisterUploadImage(app *fiber.App) {
	uploadGroup := app.Group("/upload")

	uploadGroup.Post("/", func(c *fiber.Ctx) error {
		return services.UploadImage(c)
	})
}
