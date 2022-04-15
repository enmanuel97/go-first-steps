package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expression string) int {
	valores := strings.Split(expression, "+")
	var result int

	for _, valor := range valores {
		num, error := strconv.Atoi(valor)
		if error != nil {
			fmt.Println(error)
			fmt.Println("Error: tiene que ingresar un numero entero")
			fmt.Println("o solo debes realizar con operador +")
		} else {
			result += num
		}
	}

	return result
}

func main() {
	var expression string
	var result int

	fmt.Print("Expresion => ")
	fmt.Scanln(&expression)

	result = sumar(expression)

	fmt.Printf("=> %d\n", result)
}
