package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Jesus"
	edad := 26

	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, %v y su edad es %v \n", nombre, edad)

	fmt.Print(mensaje)

	fmt.Printf("Nombre: %T \n", nombre)

	fmt.Print("Ingrese un nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Otro nombre: ", nombre)
}
