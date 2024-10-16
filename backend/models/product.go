package models

import (
    "gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Name        string  `json:"name" validate:"required"`
    Description string  `json:"description" validate:"required"`
    Price       float64 `json:"price" validate:"required"`
    Quantity    int     `json:"quantity" validate:"required"`
}

