package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InicioFuncion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello word 2!"))
}
func main() {
	//Agrupar distintas URL
	rutas := mux.NewRouter()

	//Lo que hace que cuando entremos a la pagina inicio nos va a retornar
	//una funcion
	rutas.HandleFunc("/", InicioFuncion)

	//Inizializa el puerto donde se va a initicial el servidor
	// y las rutas que va a manejar
	http.ListenAndServe(":5000", rutas)
}
