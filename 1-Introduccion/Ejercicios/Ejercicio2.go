package main

import "fmt"

func main() {
	var numero1, numero2 int
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&numero1)

	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&numero2)

	fmt.Println("El cociente es: ", numero1 / numero2)
	fmt.Println("El residuo es: ", numero1 % numero2)
}