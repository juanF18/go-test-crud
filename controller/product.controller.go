package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juanF18/go-test-crud/db"
	"github.com/juanF18/go-test-crud/models"
)

func GetProductsController(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	db.Db.Find(&products)
	json.NewEncoder(w).Encode(&products)
}
func GetProductController(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)
	db.Db.First(&product, params["id"])
	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
	}
	json.NewEncoder(w).Encode(&product)
}
func CreateProductController(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	createProduct := db.Db.Create(&product)
	err := createProduct.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&product)

}
func DeleteProductController(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)
	db.Db.First(&product, params["id"])
	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
	}
	db.Db.Delete(&product)
	w.WriteHeader(http.StatusOK)
}
