package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juanF18/go-test-crud/db"
	"github.com/juanF18/go-test-crud/models"
)

func GetCategorysController(w http.ResponseWriter, r *http.Request) {
	/*
		TODO: Controlador para ver todas las categorias
	*/
	var categories []models.Category
	db.Db.Find(&categories)
	json.NewEncoder(w).Encode(&categories)

}

func GetCategoryController(w http.ResponseWriter, r *http.Request) {
	/*
		TODO: Controlador para que solo me devuelva un usuario.
	*/
	var category models.Category
	params := mux.Vars(r)
	db.Db.First(&category, params["id"])
	if category.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	}
	json.NewEncoder(w).Encode(&category)
}

func CreateCategorysController(w http.ResponseWriter, r *http.Request) {
	/*
		TODO: Controlador para crear categorias
		Creamos una categoría y necesitamos la variable donde se almacenará.
		aca capturamos los datos que se están enviando por el body
		y lo asigna con .Decode a la variable
		para almacenar se usa el Db, el módulo de la conexión
		guardar el dato en la base de datos
		para ver si existe un error
		este es el que no devuelve un objeto
	*/
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)
	createCategory := db.Db.Create(&category)

	err := createCategory.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&category)
}
func DeleteCategorysController(w http.ResponseWriter, r *http.Request) {
	/*
		TODO: Controlador para borrar el usuario.
		El borrado de gorm usa soft delete osea que no borra el usuario
		si no que le da un fecha de eliminado.
	*/
	var category models.Category
	params := mux.Vars(r)

	db.Db.First(&category, params["id"])
	if category.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Category not found"))
		return
	}
	db.Db.Delete(&category)
	w.WriteHeader(http.StatusOK)
}
