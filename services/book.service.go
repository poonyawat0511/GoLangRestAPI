package services

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/poonyawat0511/go-fiber/models"
)

// GetBooks returns all books
func GetBooks(c *fiber.Ctx) error {
	return c.JSON(models.Books)
}

// GetBook returns a single book by ID
func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for _, book := range models.Books {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

// CreateBook adds a new book
func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	book.ID = len(models.Books) + 1
	models.Books = append(models.Books, *book)

	return c.JSON(book)
}

// UpdateBook updates an existing book by ID
func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	updatedBook := new(models.Book)
	if err := c.BodyParser(updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range models.Books {
		if book.ID == id {
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			models.Books[i] = book
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

// DeleteBook removes a book by ID
func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, book := range models.Books {
		if book.ID == id {
			models.Books = append(models.Books[:i], models.Books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
