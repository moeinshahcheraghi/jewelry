// backend/models/product.go
package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name        string  `gorm:"not null" json:"name"`
    Description string  `json:"description"`
    Price       float64 `gorm:"not null" json:"price"`
    Quantity    int     `gorm:"not null" json:"quantity"`
}

