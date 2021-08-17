package main

import "fmt"

//variable global

var i int = 23

const igv float64 = 18.0

func main() {
	// Declarar variables

	// forma tradicional
	var x string = "Hola a todos"
	fmt.Println(x)

	// Agrupaciones

	var (
		nombre string
		edad   int
	)

	nombre = "Gino"
	edad = 20

	// forma abreviada
	d := 10
	fmt.Println(d)

	nom1 := "Carlos"
	nom2 := "Juan"
	fmt.Println(nom1 == nom2)

	fmt.Println("Valor de i:", i)
	imprimir()
	fmt.Println(nombre)
	fmt.Println(edad)
}

// Las funciones en go tienden a apilarse, luego de main se apilan y se ejecutan

func imprimir() {
	fmt.Println("imprimir valor de i", i)
}
