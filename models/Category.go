package models

import "gorm.io/gorm"

/*
Modelo de datos.
*/
type Category struct {
	gorm.Model

	Name        string `gorm:"not null; unique_index"`
	Description string
	Products    []Product `gorm:"foreigKey:CategoryId"`
}
