package routes

import (
    "github.com/gin-gonic/gin"
    "jewelry/backend/controllers"
)

func PostRoutes(router *gin.Engine) {
    posts := router.Group("/api/posts")
    {
        posts.GET("/", controllers.GetPosts)
        posts.POST("/", controllers.CreatePost)
        posts.GET("/:id", controllers.GetPost)
        posts.PUT("/:id", controllers.UpdatePost)
        posts.DELETE("/:id", controllers.DeletePost)
    }
}

