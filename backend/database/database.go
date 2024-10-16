package database

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/moeinshahcheraghi/jewelry/backend/models"
)

var DB *gorm.DB

func Connect() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found. Using environment variables.")
    }

    dbHost := os.Getenv("DATABASE_HOST")
    dbUser := os.Getenv("DATABASE_USER")
    dbPassword := os.Getenv("DATABASE_PASSWORD")
    dbName := os.Getenv("DATABASE_NAME")
    dbPort := os.Getenv("DATABASE_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
        dbHost, dbUser, dbPassword, dbName, dbPort)

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Auto migrate models
    err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Story{}, &models.Complaint{}, &models.Suggestion{})
    if err != nil {
        log.Fatal("Failed to migrate database: ", err)
    }

    fmt.Println("Database connected and migrated")
}

