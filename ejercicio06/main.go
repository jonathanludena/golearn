package main

import "fmt"

func main() {
	fmt.Println(uno(5))

	// asigno en variables el resultado del retorno de la función "dos"
	numero, estado := dos(5)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println((Calculo(5, 46)))
	fmt.Println((Calculo(5, 46, 2)))
	fmt.Println((Calculo(5, 46, 10, 16)))
	fmt.Println((Calculo(5, 46, 58, 5, 7)))
	fmt.Println((Calculo(5, 46, 3, 8, 9, 12)))
}

// funcion con retorno entero
// func uno(numero int) int {
// 	return numero * 2
// }

// funcion con retorno de una variable "z" entera
func uno(numero int) (z int) {
	z = numero * 2
	return
}

// con dos valores de retorno
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// recibe parametros indefinido de enteros
func Calculo(numero ...int) int {
	total := 0
	// _ o item key de cada variable dentro de la lista; num valor del elemento en la lista
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}

	return total
}
