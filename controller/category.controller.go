package controller

import (
	"encoding/json"
	"net/http"

	"github.com/juanF18/go-test-crud/db"
	"github.com/juanF18/go-test-crud/models"
)

func GetCategorysController(w http.ResponseWriter, r *http.Request) {

}
func GetCategoryController(w http.ResponseWriter, r *http.Request) {

}
func CreateCategorysController(w http.ResponseWriter, r *http.Request) {
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

}
