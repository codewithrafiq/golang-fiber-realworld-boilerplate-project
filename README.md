## GoLang Fiber Realworld Boilerplate Project

### Instruction for setup project and creating your first app

### 1. First Clone Project

`git clone https://github.com/codewithrafiq/golang-fiber-realworld-boilerplate-project.git`

### 2. Go inside the project

`cd /golang-fiber-realworld-boilerplate-project`

### 3. Creta your First app

`./createApp`

### You have to write the app name, then it will create some folders and files, if you write 'todo' the folders and file tree is :

```
├── apps
│   └── todo
│       ├── controllers
│       │   └── todoControllers.go
│       ├── models
│       │   └── todoModel.go
│       └── routes
│           └── todoRoutes.go
├── createApp
├── go.mod
├── go.sum
├── README.md
├── server.go
└── settings
    └── database.go

```

### It will create a folder inside the apps dir, and another three folders called controllers, models, routes, and three files also, and then you have to modify 2 fills 'settings/database.go' and 'server.go'

## In 'settings/database.go' You have to input todo models and add model to the models map

`todoModel "app/apps/todo/models"`

`&todoModel.TodoModel{},`

## EX:

```
package database

import (

	"log"
	"os"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

    todoModel "app/apps/todo/models"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("project.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to database!")
	// db.Logger = logger.Default.LogMode(logger.Info) // log all queries
	log.Println("Running Migrations...")
	// Add Migrations
	models := []interface{}{
        &todoModel.TodoModel{},
	}

	db.AutoMigrate(
		models...,
	)

	Database = DbInstance{
		Db: db,
	}
}

```

## We have to modify another file

## IN 'server.go' file input todoRoutes and add todRoutes to main

`todoRoutes "app/apps/todo/routes"`

`todoRoutes.TodoRoutes(app)`

## EX :

```
package main

import (
	todoRoutes "app/apps/todo/routes"
	database "app/settings"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()

	app := fiber.New()

	// Routes
	todoRoutes.TodoRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
```

### CRUD application is ready for todo. Now if you see 'apps/todo/routes/todoRoutes.go' You can see there is Five API endpoint already created

### Ex:

```
package routes

import (
	"app/apps/todo/controllers"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app *fiber.App) {
	app.Post("/todo", controllers.CreateTodo)
	app.Get("/todo", controllers.GetTodos)
	app.Get("/todo/:id", controllers.GetTodo)
	app.Delete("/todo/:id", controllers.DeleteTodo)
	app.Put("/todo/:id", controllers.UpdateTodo)
}

```

### New if you run this project and Check those endpoints it will work

### Every time you create an app, you must modify those two files like this.

## New fill Free to Create apps you're needed

[CodeWithRafiq](https://www.linkedin.com/in/codewithrafiq)
