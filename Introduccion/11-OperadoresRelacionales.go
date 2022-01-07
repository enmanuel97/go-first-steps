package main

import "fmt"

func main() {
	a := 4
	b := 5

	var r bool

	// Igualdad
	r = a == b
	fmt.Printf("%d es igual que %d? %t \n", a, b, r)

	// Distinto
	r = a != b
	fmt.Printf("%d es igual que %d? %t \n", a, b, r)

	// Mayor que
	r = a > b
	fmt.Printf("%d es igual que %d? %t \n", a, b, r)

	// Menor que
	r = a < b
	fmt.Printf("%d es igual que %d? %t \n", a, b, r)
}