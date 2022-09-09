package config

import (
	"golangNetHttp/external/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func createRoles(db *gorm.DB) {
	roleSaved := []models.Role{}
	db.Find(&roleSaved)
	if len(roleSaved) < 2 {
		roles := []models.Role{{
			Name: "ADMIN",
		}, {
			Name: "DEV",
		}}
		db.Create(&roles)
	}
}

func runMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Language{})
}

func CreateConnection() (db *gorm.DB) {
	dsn := os.Getenv("HOST_DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	runMigration(db)
	createRoles(db)
	return db
}

func GetDB() *gorm.DB {
	return db
}
