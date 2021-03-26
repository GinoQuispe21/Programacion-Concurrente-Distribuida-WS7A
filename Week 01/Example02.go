package main

import "fmt"

func main() {
	//Declarar variables
	var largo float64 = 12.65

	ancho := 15

	name := "Gino"

	fmt.Println("El area de un rectangulo es:", largo*float64(ancho), "Y mi nombre es: ", name)
	fmt.Println("El valor del largo es mayor al valor del ancho?: ", largo > float64(ancho))
}
