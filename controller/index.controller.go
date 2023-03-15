package controller

import "net/http"

// Manejo del index
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello word 2!"))
}
