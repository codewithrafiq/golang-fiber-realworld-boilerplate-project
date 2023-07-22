package main

import (
	database "app/settings"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()

	app := fiber.New()

	// Routes

	log.Fatal(app.Listen(":3000"))
}
