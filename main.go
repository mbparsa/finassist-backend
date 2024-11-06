package main

import (
    "github.com/gin-gonic/gin"
    "github.com/mbparsa/finassist-backend/database"
    "github.com/mbparsa/finassist-backend/routes"

    // Swagger-related imports
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"

    _ "github.com/mbparsa/finassist-backend/docs" // for Swagger docs
)

// @title MyApp API
// @version 1.0
// @description API documentation for MyApp.

// @host localhost:8080
// @BasePath /api
func main() {
    // Connect to Database
    database.ConnectDB()

    // Initialize Gin Router
    r := gin.Default()

    // Setup API Routes
    routes.SetupRoutes(r)

    // Swagger URL: http://localhost:8080/swagger/index.html
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Run server
    r.Run(":8080")
}
