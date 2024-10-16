package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/moeinshahcheraghi/jewelry/backend/database"
    "github.com/moeinshahcheraghi/jewelry/backend/models"
    "github.com/moeinshahcheraghi/jewelry/backend/utils"
)

type ComplaintInput struct {
    Content string `json:"content" validate:"required"`
}

func CreateComplaint(c *gin.Context) {
    var input ComplaintInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate input
    if err := utils.Validate.Struct(input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    complaint := models.Complaint{
        Content: input.Content,
        UserID:  userID.(uint),
    }

    if err := database.DB.Create(&complaint).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create complaint"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Complaint created"})
}

func GetComplaints(c *gin.Context) {
    var complaints []models.Complaint
    if err := database.DB.Preload("User").Find(&complaints).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve complaints"})
        return
    }

    c.JSON(http.StatusOK, complaints)
}

