// backend/controllers/story_controller.go
package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "your_project_name/config"
    "your_project_name/models"
)

type CreateStoryInput struct {
    Content string `json:"content" binding:"required"`
}

func CreateStory(c *gin.Context) {
    var input CreateStoryInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    story := models.Story{
        UserID:  userID.(uint),
        Content: input.Content,
    }

    if err := config.DB.Create(&story).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create story"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": story})
}

func GetStories(c *gin.Context) {
    var stories []models.Story
    config.DB.Preload("User").Order("created_at desc").Find(&stories)
    c.JSON(http.StatusOK, gin.H{"data": stories})
}

