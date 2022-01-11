package main

import (
	"fmt"
	"strings"
)

// Closures
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	repeat3 := repeat(3)
	fmt.Println(repeat3("hola"))
	// Funcion anonima
	/* func() {
	 	fmt.Println("Hola, esta es una funcion anonima")
	 }()

	myFunc := func(nombre string) string {
		return fmt.Sprintf("Hola, %s esta es una funcion anonima", nombre)
	}

	fmt.Println(myFunc("Jesus"))*/
}
