// backend/models/story.go
package models

import "gorm.io/gorm"

type Story struct {
    gorm.Model
    UserID  uint   `json:"user_id"`
    User    User   `gorm:"foreignKey:UserID" json:"user"`
    Content string `gorm:"type:text;not null" json:"content"`
}

