package database

import (
	"SQLFiberApi/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database !\n", err.Error())
		os.Exit(2)
	}

	log.Println("Succefully connected to the database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")
	db.AutoMigrate(&models.User{})

	Database = DBInstance{DB: db}
}
