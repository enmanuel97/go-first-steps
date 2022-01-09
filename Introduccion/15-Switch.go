package main

import "fmt"

func main() {
	var vocal string
	fmt.Print("Ingrese una letra: ")
	fmt.Scanln(&vocal)

	switch {
	case vocal == "a" || vocal == "A":
		fmt.Println(vocal, "Es vocal")
	}
}