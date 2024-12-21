package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/poonyawat0511/go-fiber/services"
)

// RegisterBookRoutes ตั้งค่า route สำหรับ books
func RegisterBookRoutes(app *fiber.App, db *sql.DB) {
	bookGroup := app.Group("/books")

	bookGroup.Get("/", func(c *fiber.Ctx) error {
		return services.GetBooks(c)
	})
	bookGroup.Get("/:id", func(c *fiber.Ctx) error {
		return services.GetBook(c)
	})
	bookGroup.Post("/", func(c *fiber.Ctx) error {
		return services.CreateBook(c)
	})
	bookGroup.Put("/:id", func(c *fiber.Ctx) error {
		return services.UpdateBook(c)
	})
	bookGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return services.DeleteBook(c)
	})
}
