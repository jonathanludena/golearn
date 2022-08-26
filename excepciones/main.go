package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// ejemploDefer()
	ejemploPanic()
}

func ejemploDefer() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	/*
		la funcion no se va a ejecutar secuencialmente,
		sino cuando finalice la lectura del codigo
	*/
	defer f.Close()

	if err != nil {
		fmt.Println("error  abriendo el archivo")
		os.Exit(1)
	}
}

func ejemploPanic() {
	// se utiliza defer para tomar control o recuperarnos del error
	// a traves de una funcion anónima
	defer func() {
		reco := recover()
		// si encuentra un panic
		if reco != nil {
			log.Fatalf("ocurrió un error que generó Panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("se encontró un error!")
	}
}
