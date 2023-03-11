package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// parametros de conexion a base de datos.
var DSN = "juan_felipe:184625Juan*@tcp(127.0.0.1:3306)/demo-products-go?charset=utf8mb4&parseTime=True&loc=Local"

// Variable base de datos.
var DB gorm.DB

func DBConnection() {
	// TODO conexion a base de datos a travez del orm gorm
	var error error
	DB, error := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		if DB != nil {

			log.Println("DB connected")
		}
	}
}
