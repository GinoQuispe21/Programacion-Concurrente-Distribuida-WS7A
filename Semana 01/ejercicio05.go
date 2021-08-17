package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Ingrese su nombre: ")
	bufferIn := bufio.NewReader(os.Stdin)
	nombre, _ := bufferIn.ReadString('\n')
	nombre = strings.TrimRight(nombre, "\r\n")

	menu :=
		`
		Bienvenido, seleccione su pedido:
		[1] Pizza
		[2] Empanadas
		¿Que selecciona?
	`

	fmt.Println(menu)
	opc, _ := bufferIn.ReadString('\n')
	opc = strings.TrimRight(opc, "\r\n")

	switch opc {
	case "1":
		fmt.Println("Usted eligió Pizza")
	case "2":
		fmt.Println("Usted eligió Empanada")
	}
	fmt.Println("Gracias por su compra estimado ", nombre)
}
