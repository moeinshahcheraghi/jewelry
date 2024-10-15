// backend/controllers/product_controller.go
package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/moeinshahcheraghi/jewelry/backend/models"
    "github.com/moeinshahcheraghi/jewelry/backend/config"
)

func CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetProducts(c *gin.Context) {
    var products []models.Product
    config.DB.Find(&products)
    c.JSON(http.StatusOK, gin.H{"data": products})
}

func SearchProducts(c *gin.Context) {
    query := c.Query("q")
    var products []models.Product
    config.DB.Where("name ILIKE ?", "%"+query+"%").Find(&products)
    c.JSON(http.StatusOK, gin.H{"data": products})
}

