package routes

import (
	"github.com/gorilla/mux"
	"github.com/juanF18/go-test-crud/controller"
)

func NewRoutes() *mux.Router {
	r := mux.NewRouter()
	handleRouteHome(r)
	handleRoutesUser(r)
	return r
}

func handleRouteHome(r *mux.Router) {
	r.HandleFunc("/", controller.HomeHandler)
}

func handleRoutesUser(r *mux.Router) {
	r.HandleFunc("/users", controller.GetUsersController).Methods("GET")
	r.HandleFunc("/users/{id}", controller.GetUserController).Methods("GET")
	r.HandleFunc("/users", controller.CreateUsersController).Methods("POST")
	r.HandleFunc("/user/{id}", controller.DeleteUsersController).Methods("DELETE")

}
