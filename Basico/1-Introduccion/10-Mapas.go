package main

import "fmt"

func main() {
	dias := make(map[int]string)
	dias[0] = "Lunes"
	dias[1] = "Martes"
	dias[2] = "Miercoles"
	dias[3] = "Jueves"
	dias[4] = "Viernes"
	dias[5] = "Sabado"
	dias[6] = "Domingo"

	fmt.Println(dias)

	delete(dias, 1)

	fmt.Println(dias, len(dias))

	// Nuevo Mapa
	estudiantes := make(map[string][]int)

	estudiantes["Alex"] = []int{13, 15, 16}
	estudiantes["Joel"] = []int{14, 13, 19}

	fmt.Println(estudiantes)
	fmt.Println(estudiantes["Alex"])
}
