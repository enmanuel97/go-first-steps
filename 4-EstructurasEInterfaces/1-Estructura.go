package main

import "fmt"

// Struct Persona
type Persona struct {
	nombre string
	edad   int
}

// metodos
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editarNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

// Herencia 
type Empleado struct {
	Persona,
	sueldo float64
}

func main() {
	p1 := Persona{"Jesus Enmanuel", 24}

	p1.imprimir()
	//fmt.Println(p1)
	p1.editarNombre("Luis Manuel")

	//fmt.Println(p1)
	p1.imprimir()

	p2 := Persona{
		nombre: "Maria Isabel",
		edad:   23,
	}

	// fmt.Println(p2)
	p2.imprimir()
	p2.editarNombre("Carlos Miguel")
	p2.imprimir()

	em1 := Empleado {
		sueldo: 500,
	}
	em1.nombre = "Pedro"
	em1.edad = 35

	em1.imprimir()

}
