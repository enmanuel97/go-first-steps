package main

import "fmt"

func main() {
	var valor float64
	fmt.Print("Ingrese valor del producto: ")
	fmt.Scan(&valor)

	impuesto := valor * 0.18
	fmt.Println("El valor del producto es: ", valor)
	fmt.Println("El impuesto del producto es: ", impuesto)
	fmt.Println("El precio de venta es: ", valor + impuesto)
}
