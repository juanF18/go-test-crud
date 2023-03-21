package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juanF18/go-test-crud/db"
	"github.com/juanF18/go-test-crud/models"
)

func GetUsersController(w http.ResponseWriter, r *http.Request) {
	/*
		TODO: Me trae todos los usuario
	*/
	var users []models.User
	db.Db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.Db.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	}
	json.NewEncoder(w).Encode(&user)
}

func CreateUsersController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createUser := db.Db.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUsersController(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.Db.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.Db.Delete(&user)
	w.WriteHeader(http.StatusOK)
}
