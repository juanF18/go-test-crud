package controller

import "net/http"

// Manejo del index
/*
	Nos muestra Un hello word 2! en patanlla para saber que la ruta
	home (/) esta funcionando al iniciar el servidor
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello word 2!"))
}
