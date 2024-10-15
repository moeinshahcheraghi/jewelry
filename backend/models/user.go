// backend/models/user.go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique;not null" json:"username"`
    Password string `gorm:"not null" json:"-"`
    Email    string `gorm:"unique;not null" json:"email"`
    IsAdmin  bool   `gorm:"default:false" json:"is_admin"`
}

