package main

import "fmt"

func main() {
	exponencia(2)
}

func exponencia(numero int) {
	if numero > 100000000 {
		fmt.Println(numero)
		return
	}
	exponencia(numero * 2)
}
