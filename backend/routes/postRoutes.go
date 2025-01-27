package routes

import (
    "github.com/gin-gonic/gin"
    "jewelry/backend/controllers"
    "jewelry/backend/middlewares"  
)

func PostRoutes(router *gin.Engine) {
    posts := router.Group("/api/posts")
    {
        posts.GET("/", controllers.GetPosts)  
        posts.POST("/", middlewares.AdminRole(), controllers.CreatePost)  
        posts.GET("/:id", controllers.GetPost)
        posts.PUT("/:id", middlewares.AdminRole(), controllers.UpdatePost)  
        posts.DELETE("/:id", middlewares.AdminRole(), controllers.DeletePost) 
    }
}

