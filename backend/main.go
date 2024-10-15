// backend/main.go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/moeinshahcheraghi/jewelry/backend/config"
    "github.com/moeinshahcheraghi/jewelry/backend/models"
    "github.com/moeinshahcheraghi/jewelry/backend/routes"
)

func main() {
    config.ConnectDatabase()

    config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Story{})

    router := gin.Default()
    routes.SetupRoutes(router)
    router.Run(":8080")
}

