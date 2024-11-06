package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/mbparsa/finassit-backend/models"
    "github.com/mbparsa/finassit-backend/database"
)

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    result := database.DB.Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, gin.H{"data": users})
}
