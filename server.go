package main

import (
	"app/apps/crud/controllers"
	"app/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()

	app := fiber.New()

	// Routes
	controllers.CrudControllers(app)

	log.Fatal(app.Listen(":3000"))
}
