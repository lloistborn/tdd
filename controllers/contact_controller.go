package controllers

import (
	"github.com/gofiber/fiber/v2"
	"tdd-go-fiber/domain"
)

func GetContacts(c *fiber.Ctx, contactUseCase domain.ContactUseCase) error {
	contacts, err := contactUseCase.GetContacts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(contacts)
}
