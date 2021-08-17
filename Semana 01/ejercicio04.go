package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Ingrese un numero: ")
	// Leer datos de entrada usando buffer
	// os : paquete de donde caputra la entrada
	bufferIn := bufio.NewReader(os.Stdin) // Esta función devuelve un objeto string y un objeto error
	// Leera hasta que se realice un salto de linea
	ingreso, error := bufferIn.ReadString('\n')
	// nil == null en otros lenguajes
	if error != nil {
		fmt.Println("error: ", error.Error())
	}
	ingreso = strings.TrimRight(ingreso, "\r\n") // elimino de la derecha retorno de carro y salto de linea

	// atoi devuelve un entero y un dato error -> transforma string a int
	dato, _ := strconv.Atoi(ingreso)
	doble := dato * 2
	fmt.Println("El doble de ", dato, " es: ", doble)
}
