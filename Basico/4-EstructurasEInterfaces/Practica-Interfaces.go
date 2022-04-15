package main

import (
	"fmt"
	"math"
)

type Figura interface {
	area() float64
	perimetro() float64
}

type Cuadrado struct {
	lado float64
}

type Circulo struct {
	radio float64
}

func (cua *Cuadrado) area() float64 {
	return cua.lado * cua.lado
}

func (cua *Cuadrado) perimetro() float64 {
	return 4 * cua.lado
}

func (cir *Circulo) area() float64 {
	return math.Pi * (math.Pow(cir.radio, 2))
}

func (cir *Circulo) perimetro() float64 {
	return 2 * math.Pi * cir.radio
}

func medidas(figura Figura) {
	fmt.Println(figura)
	fmt.Println(figura.area())
	fmt.Println(figura.perimetro())
}

func main() {
	cuadrado := Cuadrado{lado: 4}
	medidas(&cuadrado)
	circulo := Circulo{radio: 5}
	medidas(&circulo)  
}
