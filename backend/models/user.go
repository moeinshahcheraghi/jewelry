package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username    string `gorm:"unique;not null" json:"username" validate:"required,min=3,max=32"`
    Email       string `gorm:"unique;not null" json:"email" validate:"required,email"`
    Password    string `gorm:"not null" json:"password" validate:"required,min=6"`
    IsAdmin     bool   `json:"is_admin"`
    Stories     []Story `json:"stories"`
    Complaints  []Complaint `json:"complaints"`
    Suggestions []Suggestion `json:"suggestions"`
}

