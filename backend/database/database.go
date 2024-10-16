package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"

    "github.com/moeinshahcheraghi/jewelry/backend/models"
)

var DB *gorm.DB

func Connect() {
    var err error
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        // Example DSN, replace with your actual database credentials
        dsn = "host=localhost user=postgres password=postgres dbname=jewelry port=5432 sslmode=disable TimeZone=Asia/Tehran"
    }
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Auto migrate models
    err = DB.AutoMigrate(&models.User{}, &models.Story{}, &models.Complaint{}, &models.Suggestion{})
    if err != nil {
        log.Fatal("Failed to migrate database: ", err)
    }

    fmt.Println("Database connected and migrated")
}

