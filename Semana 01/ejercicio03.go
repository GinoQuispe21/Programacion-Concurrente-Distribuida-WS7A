package main

import "fmt"

func main() {
	fmt.Println("Ingrese un numero: ")

	// Leer input mediante fmt.Scanf -> "%" para el formato de flaot - &num para guardar el valor en la direccion de memoria de la variable (punteros)
	var num float64
	fmt.Scanf("%f", &num)

	doble := num * 2
	fmt.Println("El doble del valor ingresado es: ", doble)
}
