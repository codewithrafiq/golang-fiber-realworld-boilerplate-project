package services

import (
	"app/apps/crud/entities"
	"app/database"

	"github.com/gofiber/fiber/v2"
)

func CreateCrud(c *fiber.Ctx) error {
	var crud entities.CrudEntity

	if err := c.BodyParser(&crud); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"success": false,
			"message": "Unable to parse JSON",
		})

	}

	database.Database.Db.Create(&crud)

	crudResponse := entities.ResponseCrud(crud)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    crudResponse,
	})
}

func GetCruds(c *fiber.Ctx) error {
	var crud []entities.CrudEntity

	database.Database.Db.Find(&crud)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    crud,
	})
}
