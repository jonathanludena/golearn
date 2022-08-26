package main

import "fmt"

var estado bool

func main() {
	estado = true
	if estado = false; estado == true { // se permite asignar valores en variables mientras se realiza la condicion
		fmt.Println(estado)
	} else {
		fmt.Println("Es falso")
	}

	switch numero := 10; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mayor a 5")
	}

}
