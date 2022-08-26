package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readFile()
	println("")
	readFile2()
	saveFile()
	saveFile2()
}

func saveFile() {
	file, err := os.Create("./nuevo_archivo.txt")

	if err != nil {
		fmt.Println("Hubo un error")
		return
	}

	fmt.Fprintln(file, "Esta es una prueba de archivo nuevo")
	file.Close()
}

func saveFile2() {
	fileName := "./nuevo_archivo.txt"
	if Append(fileName, "\nEsta es una segunda linea") == false {
		fmt.Println("Error en la segunda linea")
	}
}

func Append(file string, texto string) bool {
	file_other, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	_, err = file_other.WriteString(texto)

	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	return true
}

// Si existe archivo.txt va a leer el archivo
// y se asigna a file, sino asigna true en err

func readFile() {
	file, err := ioutil.ReadFile("./archivo.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(file))
	}
}

func readFile2() {
	file, err := os.Open("./archivo.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			register := scanner.Text()
			fmt.Printf("Linea >" + register + "\n")
		}
	}

	file.Close()
}
