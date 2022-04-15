package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {
	// Relacion Uno a Uno
	// alex := User{
	// 	nombre: "Alex",
	// 	email:  "alex@gmail.com",
	// 	activo: true,
	// }
	// roel := User{
	// 	nombre: "Roel",
	// 	email:  "roel@gmail.com",
	// 	activo: true,
	// }

	// alexStudent := Student {
	// 	user: alex,
	// 	codigo: "00125",
	// }

	// fmt.Println(roel)
	// fmt.Println(alexStudent)

	// Relacion Uno a Muchos
	video1 := Video{titulo: "01-Introduccion"}
	video2 := Video{titulo: "02-Instalacion"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	fmt.Println(video1.curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
