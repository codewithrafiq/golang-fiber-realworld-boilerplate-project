package controllers

import (
	services "app/apps/crud/services"

	"github.com/gofiber/fiber/v2"
)

func CrudControllers(app *fiber.App) {
	app.Post("/crud", services.CreateCrud)
	app.Get("/crud", services.GetCruds)
}
