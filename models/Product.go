package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string  `gorm:"not null;unique_index" json:"name"`
	Description string  `json:"description"`
	Price       float32 `gorm:"not null" json:"price"`
	Amount      uint16  `gorm:"not null" json:"amount"`
	CategoryId  uint
}
