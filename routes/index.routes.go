package routes

import (
	"github.com/gorilla/mux"
	"github.com/juanF18/go-test-crud/controller"
)

func NewRoutes() *mux.Router {
	r := mux.NewRouter()
	handleRouteHome(r)
	handleRoutesUser(r)
	handleRouteProducts(r)
	handleRouteCategory(r)
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

func handleRouteProducts(r *mux.Router) {
	r.HandleFunc("/producs", controller.GetProductsController).Methods("GET")
	r.HandleFunc("/producs/{id}", controller.GetProductController).Methods("GET")
	r.HandleFunc("/producs", controller.CreateProductController).Methods("POST")
	r.HandleFunc("/producs/{id}", controller.DeleteProductController).Methods("DELETE")
}

func handleRouteCategory(r *mux.Router) {
	r.HandleFunc("/category", controller.GetCategorysController).Methods("GET")
	r.HandleFunc("/category/{id}", controller.GetCategoryController).Methods("GET")
	r.HandleFunc("/category", controller.CreateCategorysController).Methods("POST")
	r.HandleFunc("/category/{id}", controller.DeleteCategorysController).Methods("DELETE")
}
