// backend/routes/routes.go
package routes

import (
    "your_project_name/controllers"
    "your_project_name/middleware"

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

