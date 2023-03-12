package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string `gorm:"not null;unique_index"`
	Description string
	Price       float32 `gorm:"not null"`
	Amount      uint16  `gorm:"not null"`
	Categoryid  uint
}
