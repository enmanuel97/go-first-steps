package main

import "fmt"

// Struct Persona
type Persona struct {
	nombre string
	edad   int
}

// metodos
func (p *Persona) imprimir() {
	
}

func main() {
	p1 := Persona{"Jesus Enmanuel", 24}
	
	fmt.Println(p1)
	p1.nombre = "Luis Miguel"

	fmt.Println(p1)

	p2 := Persona {
		nombre: "Maria Isabel",
		edad: 23,
	}

	fmt.Println(p2)
}
