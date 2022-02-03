package main

import (
	"fmt"
	"strconv"
)

var numero int // variable privada dentro del package main
var texto string
var status bool

var Numero int // variable que puede ser exportada fuera del package

// float32 float64
// int8 int16
// enteros sin signo uint

func main() {
	var numero2 int
	numero2 = 0
	numero3 := 4 // no hace falta declarar que es variable de tipo entera. Y solo se puede usar una sola vez por variable (":=")
	numero3 = 15

	var numero4, numero5, numero6 int // se pueden declarar varias variables en cadena
	numero4 = 1
	numero5 = 2
	numero6 = 3
	numero7, numero8, texto, status := 2, 5, "Esto es un texto", true // declarar sin var variables con asignación en cadena

	var moneda int64 = 0
	numero2 = int(moneda)

	texto = fmt.Sprintf("%d", moneda) // convierte con el verbo %d decimal a un string
	texto = strconv.Itoa(int(moneda)) // Itoa solo acepta solo var int puro

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)
	fmt.Println(numero6)
	fmt.Println(numero7)
	fmt.Println(numero8)

	fmt.Println(texto)
	fmt.Println(status) // los bool se asignan por default a false

	MostrarStatus()
}

func MostrarStatus() {
	fmt.Println(status)
}
