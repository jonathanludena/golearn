package main

import "fmt"

var matriz [5][7]int

func main() {
	// tabla := [10]int{5, 6, 7, 8, 12, 2, 3, 5, 11, 12}

	// for i := 0; i < len(tabla); i++ {
	// 	fmt.Println((tabla[i]))
	// }

	matriz[3][5] = 1
	fmt.Println((matriz))
}
