package database

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    dsn := "host=localhost user=youruser password=yourpassword dbname=myapp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    fmt.Println("Database connected successfully.")
}
