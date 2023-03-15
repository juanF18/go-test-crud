package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// parametros de conexion a base de datos.
var DSN = ConfigDb.user + ":" + ConfigDb.password + "@tcp(" + ConfigDb.ip + ")/" + ConfigDb.name_db + "?charset=utf8mb4&parseTime=True&loc=Local"
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
