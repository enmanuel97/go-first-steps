package main

import "fmt"

func main() {
	nombres := [...]string{"Jesus", "Luis", "Miguel"}

	// for i := 0; i < len(nombres); i++ {
	// 	fmt.Println(nombres[i])
	// }

	for _, elemento := range nombres {
		fmt.Println(elemento)
	}
}
