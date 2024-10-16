package main

import (
    "github.com/gin-gonic/gin"
    "github.com/moeinshahcheraghi/jewelry/backend/controllers"
    "github.com/moeinshahcheraghi/jewelry/backend/database"
    "github.com/moeinshahcheraghi/jewelry/backend/middleware"
)

func main() {
    database.Connect()
    router := gin.Default()

    // Public routes
    router.POST("/api/register", controllers.Register)
    router.POST("/api/login", controllers.Login)

    // Protected routes
    protected := router.Group("/api")
    protected.Use(middleware.AuthMiddleware())
    {
        // Admin routes
        protected.GET("/users", controllers.GetUsers)
        protected.PUT("/users/:id/promote", controllers.PromoteToAdmin)
        protected.DELETE("/users/:id", controllers.DeleteUser)

        // Stories
        protected.POST("/stories", controllers.CreateStory)
        protected.GET("/stories", controllers.GetStories)

        // Complaints
        protected.POST("/complaints", controllers.CreateComplaint)
        protected.GET("/complaints", controllers.GetComplaints)

        // Suggestions
        protected.POST("/suggestions", controllers.CreateSuggestion)
        protected.GET("/suggestions", controllers.GetSuggestions)

        // Products
        protected.POST("/products", controllers.CreateProduct)
        protected.GET("/products", controllers.GetProducts)

        // Search
        protected.GET("/search", controllers.Search)
    }

    router.Run(":8080")
}

