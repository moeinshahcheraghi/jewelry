package models

import (
    "gorm.io/gorm"
)

type Complaint struct {
    gorm.Model
    Content string `gorm:"type:text;not null" json:"content" validate:"required"`
    UserID  uint   `json:"user_id"`
    User    User   `json:"user"`
}

