package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/moeinshahcheraghi/jewelry/backend/database"
    "github.com/moeinshahcheraghi/jewelry/backend/models"
    "github.com/moeinshahcheraghi/jewelry/backend/utils"
)

type SuggestionInput struct {
    Content string `json:"content" validate:"required"`
}

func CreateSuggestion(c *gin.Context) {
    var input SuggestionInput
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

    suggestion := models.Suggestion{
        Content: input.Content,
        UserID:  userID.(uint),
    }

    if err := database.DB.Create(&suggestion).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create suggestion"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Suggestion created"})
}

func GetSuggestions(c *gin.Context) {
    var suggestions []models.Suggestion
    if err := database.DB.Preload("User").Find(&suggestions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve suggestions"})
        return
    }

    c.JSON(http.StatusOK, suggestions)
}

