package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/poonyawat0511/go-fiber/services"
)

// RegisterUserRoutes ลงทะเบียนเส้นทางสำหรับ user
func RegisterUserRoutes(app *fiber.App, db *sql.DB) {
	userGroup := app.Group("/users")

	userGroup.Post("/", func(c *fiber.Ctx) error {
		return services.CreateUserHandle(c, db)
	})
	userGroup.Get("/", func(c *fiber.Ctx) error {
		return services.GetUsersHandle(c, db)
	})
	userGroup.Get("/:id", func(c *fiber.Ctx) error {
		return services.GetUserHandle(c, db)
	})
	userGroup.Put("/:id", func(c *fiber.Ctx) error {
		return services.UpdateUserHandle(c, db)
	})
	userGroup.Delete("/:id", func(c *fiber.Ctx) error {
		return services.DeleteUserHandle(c, db)
	})
}