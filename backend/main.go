// backend/main.go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/yourusername/my-ecommerce-app/backend/config"
    "github.com/yourusername/my-ecommerce-app/backend/models"
    "github.com/yourusername/my-ecommerce-app/backend/routes"
)

func main() {
    config.ConnectDatabase()

    config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Story{})

    router := gin.Default()
    routes.SetupRoutes(router)
    router.Run(":8080")
}

