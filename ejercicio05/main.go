package main

import "fmt"

func main() {

	var i int = 0

RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("voy a RUTINA  sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("valor de i: %d\n", i)
		i++
	}

	// var i = 0
	// for i < 10 {
	// 	fmt.Printf("\n Valor de i: %d", i)
	// 	if i == 5 {
	// 		fmt.Printf(" multiplicamos por 2 \n")
	// 		i = i * 2
	// 		continue // sigue en la iteracion sin pasar por las siguientes lineas
	// 	}
	// 	fmt.Printf("     Paso por este lugar \n")
	// 	i++
	// }

	// numero := 0
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Continuo")
	// 	fmt.Println("Ingrese un número secreto")
	// 	fmt.Scanln(&numero)
	// 	if numero == 100 {
	// 		break
	// 	}
	// }
}
