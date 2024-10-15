// backend/routes/routes.go
package routes

import (
    "github.com/moeinshahcheraghi/jewelry/backend/controllers"
    "github.com/moeinshahcheraghi/jewelry/backend/middleware"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/register", controllers.Register)
    router.POST("/login", controllers.Login)
    router.GET("/products", controllers.GetProducts)
    router.GET("/search", controllers.SearchProducts)

    admin := router.Group("/admin")
    admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
    admin.POST("/products", controllers.CreateProduct)

    stories := router.Group("/stories")
    stories.Use(middleware.AuthMiddleware())
    stories.POST("/", controllers.CreateStory)
    router.GET("/stories", controllers.GetStories)
}

