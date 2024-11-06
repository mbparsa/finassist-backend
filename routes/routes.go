package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/mbparsa/finassist-backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.POST("/users", controllers.CreateUser)
        api.GET("/users", controllers.GetUsers)
    }
}
