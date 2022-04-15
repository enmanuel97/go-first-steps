package main

import "fmt"

func main() {

	// Slicen
	numeros := []int{1,2,3}

	fmt.Println(numeros, len(numeros))

	// Agregar datos a un slicen
	nuevoSlicen := append(numeros, 4)

	fmt.Println(nuevoSlicen, len(numeros))

	// Sub Slicen
	subSlice := numeros[:2]

	numeros[0] = 100

	fmt.Println(subSlice)
	fmt.Println(numeros)

	// Punteros
	// Longitud
	// Capacidad

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
}