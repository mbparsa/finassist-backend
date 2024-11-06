package database

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"github.com/mbparsa/finassist-backend/models"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
	dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=finassist port=5432 sslmode=disable TimeZone=America/Los_Angeles", dbUsername, dbPassword)
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
