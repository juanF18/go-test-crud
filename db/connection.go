package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// parametros de conexion a base de datos.
var DSN = "root:184625Topo*@tcp(127.0.0.1:3306)/demo-products-go?charset=utf8mb4&parseTime=True&loc=Local"
var Db *gorm.DB

// Variable base de datos.
// Tiene metodos para usar en base de datos.

func DBConnection() {
	// TODO conexion a base de datos a travez del orm gorm
	var err error
	Db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

}
