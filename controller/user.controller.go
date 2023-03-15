package controller

import "net/http"

func GetUsersController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUsersHandler"))
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserHandler"))
}

func CreateUsersController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUserHandler"))
}

func DeleteUsersController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUserHandler"))
}
