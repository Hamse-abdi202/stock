package infra

import (
	"fmt"
	"log"

	"Github.com/Hams/abdi/InventoryManagmentSystem/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	config := Configurations

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", 
	config.DBHost, config.DBUser, config.DBPassword, config.DBPort, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// postgres.Open()
	if err != nil {
		log.Fatal("❌ failed to connect database")
	}

	if err := db.AutoMigrate(
		models.User{},
		models.Category{},
		models.Product{},
		models.StockMovement{},
	); err != nil {
		log.Fatal("❌ Failed to auto migrate the database")
	}

	DB = db

}
