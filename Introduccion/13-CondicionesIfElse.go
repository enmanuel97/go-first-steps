package main

import "fmt"

func main() {
	/* App para restaurante
	Descuentos de 10% hasta 100 de consumo
	Descuentos de 20% hasta 200 de consumo
	Descuentos de 30% mayor 200 de consumo
	*/
	var consumo, descuento float64
	var datosDescuento string

	// Entrada de datos
	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		if consumo <= 100 {
			datosDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 {
			datosDescuento = "20%"
			descuento = consumo * 0.20
		} else {
			datosDescuento = "30%"
			descuento = consumo * 0.30
		}
	} else {
		fmt.Println("Error al ingresar el consumo")
	}

	montoDescuento := consumo - descuento
	itbis := montoDescuento * 0.18

	totalAPagar := montoDescuento + itbis

	fmt.Println("-------Factura de consumo-------")
	fmt.Println("Descuento que aplica: ", datosDescuento)
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Monto con descuento: ", descuento)
	fmt.Println("Descuento que aplica: ", montoDescuento)
	fmt.Println("ITBIS: ", itbis)
	fmt.Println("Total a pagar: ", totalAPagar)
}
