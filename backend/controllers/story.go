package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/moeinshahcheraghi/jewelry/backend/database"
    "github.com/moeinshahcheraghi/jewelry/backend/models"
    "github.com/moeinshahcheraghi/jewelry/backend/utils"
)

type StoryInput struct {
    Content string `json:"content" validate:"required"`
}

func CreateStory(c *gin.Context) {
    var input StoryInput
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

    story := models.Story{
        Content: input.Content,
        UserID:  userID.(uint),
    }

    if err := database.DB.Create(&story).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create story"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Story created"})
}

func GetStories(c *gin.Context) {
    var stories []models.Story
    if err := database.DB.Preload("User").Find(&stories).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stories"})
        return
    }

    c.JSON(http.StatusOK, stories)
}

