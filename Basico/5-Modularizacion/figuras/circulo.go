package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (cir *Circulo) area() float64 {
	return math.Pi * (math.Pow(cir.Radio, 2))
}

func (cir *Circulo) perimetro() float64 {
	return 2 * math.Pi * cir.Radio
}
