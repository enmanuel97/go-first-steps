package main

import "net/http"

func main() {
	// Crear Servidor
	http.ListenAndServe("Localhost:3000", nil)
}
