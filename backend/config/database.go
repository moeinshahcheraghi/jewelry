package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
        return nil, err
    }

    log.Println("Database connection established successfully.")
    return DB, nil
}

