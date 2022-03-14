package main

<<<<<<< HEAD
import (
	"fmt"
	"time"

	us "./user"
)

// Herencia de user
type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Pablo S", time.Now(), true)
	fmt.Println((u.Usuario))
=======
import "fmt"

func main() {
	// con func make => asignamos a paises un map con clave string y valor de tipo string y le doy un tamaño de 5
	paises := make(map[string]string, 5)
	// fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	// fmt.Println(paises["Mexico"])
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	// Agregar elementos al map
	campeonato["River Plate"] = 25
	// Modificar elementos al map
	campeonato["Chivas"] = 55
	// Eleminiar elementos del map
	delete(campeonato, "Real Madrid")

	fmt.Println(campeonato)

	// Listando todos los elementos del map
	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	}

	// Buscando elemento
	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado %d, y el equipo existe %t\n", puntaje, existe)

	puntaje, existe = campeonato["Chivas"]
	fmt.Printf("El puntaje capturado %d, y el equipo existe %t\n", puntaje, existe)
>>>>>>> a269923d8dd6b9ae2824092a591f1b930ec14e87
}
