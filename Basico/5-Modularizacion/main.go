package main

// import "github.com/donvito/hellomod"
import "github.com/enmanuel97/figuras"

func main() {

	// hellomod.SayHello()

	cua := figuras.Cuadrado{Lado: 8}
	figuras.Medidas(&cua)

	cir := figuras.Circulo{Radio: 5}
	figuras.Medidas(&cir)
	/*p1 := models.Persona{}
	p1.Constructor("Jesus", 26)
	p1.SetEdad(24)
	p1.SetNombre("Jesus Enmanuel")
	fmt.Println(p1)*/
}
