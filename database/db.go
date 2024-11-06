package database

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"github.com/mbparsa/finassist-backend/models"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    dsn := "host=localhost user=postgres password=<>! dbname=finassist port=5432 sslmode=disable TimeZone=America/Los_Angeles"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    fmt.Println("Database connected successfully.")

	// Automatically create the "users" table based on the User model
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	fmt.Println("Database migrated successfully.")

}
