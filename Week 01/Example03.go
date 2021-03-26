package main

import "fmt"

func main() {

	//Declaramos un array

	array_int := [3]int{20, 21, 22}

	for i := 0; i < 3; i++ {
		println(array_int[i])
	}

	println(len(array_int))

	for i, v := range array_int {
		fmt.Printf("El valor en v es %d en la posicion %d\n", v, i)
	}

	for _, v := range array_int {
		fmt.Printf("\n\nEl valor de v es %d\n", v)
	}
	i := 0
	for {
		if i == 3 {
			break
		}
		fmt.Println("Numero de intentos = ", i)
		i++
	}

	j := 0

	for i > 0 {
		if j == 5 {
			break
		}
		j++
	}
	fmt.Println("Numero de intentos de J = ", j)
}
