package books

import (
	"github.com/Amarendrapratihari/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateBookRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Price = body.Price
	book.Description = body.Description

	// save book
	h.DB.Save(&book)

	return c.Status(fiber.StatusOK).JSON(&book)
}
