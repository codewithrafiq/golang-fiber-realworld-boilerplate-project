package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	models := []interface{}{}

	db.AutoMigrate(
		models...,
	)

	Database = DbInstance{
		Db: db,
	}
}
