package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func createModel(appName string) {
	text := `
package models

import "time"

type ` + firstCharToUpper(appName) + `Model struct {
	ID        uint   ` + "`json:\"id\" gorm:\"primary_key\"`" + `
	Title     string ` + "`json:\"title\"`" + `
	CreatedAt time.Time
}
`

	createFile(appName, "models", appName+"Model.go", text)
}

func createController(appName string) {
	text := `
package controllers

import (
	model "app/apps/` + appName + `/models"
	database "app/settings"

	"github.com/gofiber/fiber/v2"
)

func Create` + firstCharToUpper(appName) + `(c *fiber.Ctx) error {
	var crud model.` + firstCharToUpper(appName) + `Model

	if err := c.BodyParser(&crud); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"success": false,
			"message": "Unable to parse JSON",
		})
	}

	database.Database.Db.Create(&crud)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    crud,
	})
}

func Get` + firstCharToUpper(appName) + `s(c *fiber.Ctx) error {
	var crud []model.` + firstCharToUpper(appName) + `Model

	database.Database.Db.Find(&crud)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    crud,
	})
}

func Delete` + firstCharToUpper(appName) + `(c *fiber.Ctx) error {
	id := c.Params("id")

	var crud model.` + firstCharToUpper(appName) + `Model

	database.Database.Db.Find(&crud, id)

	if crud.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "No ` + strings.ToLower(appName) + ` found with ID",
		})
	}

	database.Database.Db.Delete(&crud)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Deleted Successfully !",
	})
}

func Update` + firstCharToUpper(appName) + `(c *fiber.Ctx) error {
	id := c.Params("id")

	var crud model.` + firstCharToUpper(appName) + `Model

	database.Database.Db.Find(&crud, id)

	if crud.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "No ` + strings.ToLower(appName) + ` found with ID",
		})
	}

	if err := c.BodyParser(&crud); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"success": false,
			"message": "Unable to parse JSON",
		})
	}

	database.Database.Db.Save(&crud)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    crud,
	})
}

func Get` + firstCharToUpper(appName) + `(c *fiber.Ctx) error {
	id := c.Params("id")

	var crud model.` + firstCharToUpper(appName) + `Model

	database.Database.Db.Find(&crud, id)

	if crud.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "No ` + strings.ToLower(appName) + ` found with ID",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    crud,
	})
}




`

	createFile(appName, "controllers", appName+"Controllers.go", text)
}

func createRoute(appName string) {
	text := `package routes

import (
	"app/apps/` + appName + `/controllers"

	"github.com/gofiber/fiber/v2"
)

func ` + firstCharToUpper(appName) + `Routes(app *fiber.App) {
	app.Post("/` + strings.ToLower(appName) + `", controllers.Create` + firstCharToUpper(appName) + `)
	app.Get("/` + strings.ToLower(appName) + `", controllers.Get` + firstCharToUpper(appName) + `s)
	app.Get("/` + strings.ToLower(appName) + `/:id", controllers.Get` + firstCharToUpper(appName) + `)
	app.Delete("/` + strings.ToLower(appName) + `/:id", controllers.Delete` + firstCharToUpper(appName) + `)
	app.Put("/` + strings.ToLower(appName) + `/:id", controllers.Update` + firstCharToUpper(appName) + `)
}
`

	createFile(appName, "routes", appName+"Routes.go", text)
}

func createFile(appName, dirName, fileName, content string) {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	filePath := filepath.Join(directory, "apps", appName, dirName, fileName)

	err = os.MkdirAll(filepath.Join(directory, "apps", appName, dirName), 0755)
	if err != nil {
		fmt.Println("Error creating directory path:", err)
		return
	}

	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File", filePath, "written successfully.")
}

func firstCharToUpper(s string) string {
	if len(s) < 1 {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
func main() {
	fmt.Print("Enter your app name: ")
	var appName string
	fmt.Scanln(&appName)

	// Remove invalid characters from the app name using a regular expression
	re := regexp.MustCompile(`[\/:*?"<>|]`)
	appName = re.ReplaceAllString(appName, "")

	createModel(appName)
	createController(appName)
	createRoute(appName)
}
