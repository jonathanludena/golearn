package main

import (
	"fmt"
	"time"
)

func main() {
	// creamos el canal a traves de make con typo chan y el canal sera de tipo time.Duration
	// con time.Duration se graba el calculo de la diferencia entre diferentes horas
	canal1 := make(chan time.Duration)

	// ejecutamos una gorutine con funcion bucle
	go bucle(canal1)

	fmt.Println("Llegué hasta aquí")

	// a la variable msg le vamos a asignar el resultado de
	// lo obtuvo canal1 cuando se ejecuto bucle
	msg := <-canal1
	// en cuanto lo reciba msg podrá mostrar su valor
	// el valor esperado es el tiempo que le tomo al
	// computador ejecutar bucle y su ciclo for
	fmt.Println(msg)
}

func bucle(canal chan time.Duration) {
	// inicio contiene fecha/hora al inicio de
	// func bucle
	inicio := time.Now()
	for i := 0; i < 10000000000; i++ {
		// no realiza nada el bucle pero
		// tomará un tiempo en terminar el bucle for
	}
	// asignamos la fecha/hora en el momento que termino
	// de ejecutarse el bucle for
	final := time.Now()
	// enviamos el resultado a través de Sub quien hace
	// el calculo de la diferencia de tiempo inicial y final
	// asignamos por medio del canal el valor a la gorutine
	canal <- final.Sub(inicio)
}
