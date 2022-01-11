package main

import "fmt"

func main() {
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)

	// Array con valores
	nombres := [3]string {"Jesus", "Enmanuel", "De La Cruz"}

	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Verde",
		"Negro",
		"Azul",
	}

	fmt.Println(colores, len(colores))
}
