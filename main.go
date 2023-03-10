package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juanF18/go-test-crud/db"
	"github.com/juanF18/go-test-crud/models"
	"github.com/juanF18/go-test-crud/routes"
)

func main() {
	//Conexion a base de datos
	db.DBConnection()

	err := db.Db.AutoMigrate(models.Category{}, models.Product{}, models.User{})
	if err != nil {
		log.Println("Error al migrar")
	}

	//Agrupar distintas URL
	rutas := mux.NewRouter()

	//Lo que hace que cuando entremos a la pagina inicio nos va a retornar
	//una funcion
	rutas.HandleFunc("/", routes.HomeHandler)

	//Inizializa el puerto donde se va a initicial el servidor
	// y las rutas que va a manejar
	http.ListenAndServe(":5000", rutas)
}
