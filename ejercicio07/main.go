package main

import "fmt"

// Funciones anónimas

//  tipo de variable funcion
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma 5 + 7 es = %d \n", Calculo(5, 7))

	// Restamos - redifinimos la funcion anónima - cambiamos la funcion
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Resta 6 - 4 es = %d \n", Calculo(6, 4))

	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Divide 12 / 3 es = %d \n", Calculo(12, 3))
	Operaciones()

	/* CLOSURE */
	TablaDel := 2
	MiTabla := Tabla(TablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13

		return a * b

	}
	fmt.Println(resultado())
}

// Closures
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
