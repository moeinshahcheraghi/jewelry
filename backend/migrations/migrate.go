package migrations

import (
    "jewelry/backend/models"
    "gorm.io/gorm"
    "log"
)

func Migrate(db *gorm.DB) error {
    // اجرای Migration
    err := db.AutoMigrate(&models.Post{})
    if err != nil {
        log.Printf("Migration failed: %v", err)
        return err
    }
    log.Println("Migration completed successfully.")
    return nil
}

