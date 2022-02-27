package main

import "fmt"

// var matriz [5][7]int

func main() {
	// tabla := [10]int{5, 6, 7, 8, 12, 2, 3, 5, 11, 12}

	// for i := 0; i < len(tabla); i++ {
	// 	fmt.Println((tabla[i]))
	// }

	// matriz[3][5] = 1
	// fmt.Println((matriz))

	// Slices
	// matriz := []int{2, 5, 4}
	// fmt.Println(matriz)

	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]   // desde el tercer elemento hasta el final del vector
	porcion2 := elementos[:3]  // desde el primer elemento hasta el tercer elemento del vector
	porcion3 := elementos[2:4] // desde el primer elemento hasta el tercer elemento del vector

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func variante3() {
	// make => tiene 3 parametros para crear slices: type, len, capacity
	elementos := make([]int, 5, 20)
	// cap => func que permite revisar la capacidad de un vector
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
}
