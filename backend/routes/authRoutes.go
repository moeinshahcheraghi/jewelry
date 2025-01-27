package routes

import (
    "github.com/gin-gonic/gin"
    "jewelry/backend/controllers"
)

func AuthRoutes(router *gin.Engine) {
    auth := router.Group("/auth")
    {
        auth.POST("/register", controllers.Register)
        auth.POST("/login", controllers.Login)
    }
}

