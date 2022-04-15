package figuras

import "fmt"

type Figura interface {
	area() float64
	perimetro() float64
}

func Medidas(figura Figura) {
	fmt.Println(figura)
	fmt.Println("Area:", figura.area())
	fmt.Println("Perimetro:", figura.perimetro())
}
