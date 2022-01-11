package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups! al parecer hubo un error")
		}
	}()

	if file, error := os.Open("hol.txt"); error != nil {
		panic("No fue posible obtener el archivo!!")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo!")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])

		fmt.Println(contenidoArchivo)
	}
}
