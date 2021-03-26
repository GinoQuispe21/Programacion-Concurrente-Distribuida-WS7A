package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Bloque textual multilinea
	menu := `
			Bienvenido
			[1] Pizza
			[2] Empanada
			¿Que prefiere?
	`

	fmt.Println(menu)

	//Lectura de datos
	bufferIn := bufio.NewReader(os.Stdin)
	datoIn, _ := bufferIn.ReadString('\n')
	dato := strings.Trim(datoIn, "\r\n")

	switch dato {
	case "1":
		fmt.Println("Elegiste pizza")
	case "2":
		fmt.Println("Elegiste empanada")
	default:
		fmt.Println("Opcion Invalida")
	}

}
