package books

import (
	"github.com/Amarendrapratihari/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func (h handler) AddBook(c *fiber.Ctx) error {
	body := AddBookRequestBody{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Price = body.Price
	book.Description = body.Description

	// insert new db entry
	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&book)

}
