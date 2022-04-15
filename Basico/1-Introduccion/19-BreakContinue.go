package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {

		if i == 5 {
			fmt.Println("Salta la interacion")
			continue
		}

		if i == 8 {
			fmt.Println("Romper con bucle")
			break
		}

		fmt.Println(i)
	}
}
