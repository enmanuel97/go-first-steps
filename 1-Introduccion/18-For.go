package main

import "fmt"

func main() {
	// Bucle Infinito
	// for {
	// 	fmt.Println("Hola Mundo")
	// }

	// Bucle tipo while
	numeros := 12455
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}

	fmt.Println("La cantidad de digitos es: ", c)

	// Bucle for
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("El numero es: ", i)
		}
	}
}
