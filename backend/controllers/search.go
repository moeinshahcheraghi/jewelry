package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/moeinshahcheraghi/jewelry/backend/database"
    "github.com/moeinshahcheraghi/jewelry/backend/models"
)

func Search(c *gin.Context) {
    query := c.Query("q")
    if query == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
        return
    }

    var products []models.Product
    var stories []models.Story

    // Search in products
    database.DB.Where("name ILIKE ?", "%"+query+"%").Find(&products)

    // Search in stories
    database.DB.Where("content ILIKE ?", "%"+query+"%").Find(&stories)

    c.JSON(http.StatusOK, gin.H{
        "products": products,
        "stories":  stories,
    })
}

