package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola,", nombre)
}

func sumar(a, b int) int {
	return a + b
}

func main() {
	saludar("Jesus")

	fmt.Println("La sumatoria es:", sumar(2, 5))
}
