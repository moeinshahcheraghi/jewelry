package main

import (
    "log"
    "os"
    "jewelry/backend/config"
    "jewelry/backend/routes"
    "jewelry/backend/migrations"  
    "github.com/gin-gonic/gin"
)

func main() {
    db, err := config.ConnectDatabase()  
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    if err := migrations.Migrate(db); err != nil { 
        log.Fatalf("Migration failed: %v", err)
    }

    r := gin.Default()

    routes.AuthRoutes(r)

    routes.PostRoutes(r)

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }

    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

