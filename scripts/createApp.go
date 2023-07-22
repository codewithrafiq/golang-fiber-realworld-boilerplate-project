package main

import (
	"fmt"
	"os"
	"regexp"
)

func createControllerFile(appName string) {
	text := `package controllers`

	directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	filePath := fmt.Sprintf("%s\\apps\\%s\\controllers\\%sController.go", directory, appName, appName)

	err = os.MkdirAll(directory+"\\apps\\"+appName+"\\controllers", 0755)
	if err != nil {
		fmt.Println("Error creating directory path:", err)
		return
	}

	err = os.WriteFile(filePath, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File", filePath, "written successfully.")
}

func createServiceFile(appName string) {
	text := `package services`

	directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}
	filePath := fmt.Sprintf("%s\\apps\\%s\\services\\%sService.go", directory, appName, appName)

	err = os.MkdirAll(directory+"\\apps\\"+appName+"\\services", 0755)
	if err != nil {
		fmt.Println("Error creating directory path:", err)
		return
	}
	err = os.WriteFile(filePath, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func createEntityFile(appName string) {
	text := `package entities`

	directory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}
	filePath := fmt.Sprintf("%s\\apps\\%s\\entities\\%sEntity.go", directory, appName, appName)

	err = os.MkdirAll(directory+"\\apps\\"+appName+"\\entities", 0755)
	if err != nil {
		fmt.Println("Error creating directory path:", err)
		return
	}
	err = os.WriteFile(filePath, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	fmt.Print("Enter your app name: ")
	var appName string
	fmt.Scanln(&appName)
	// Remove invalid characters from the app name using a regular expression
	re := regexp.MustCompile(`[\/:*?"<>|]`)
	appName = re.ReplaceAllString(appName, "")
	createControllerFile(appName)
	createServiceFile(appName)
	createEntityFile(appName)
}
