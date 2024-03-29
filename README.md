# Conceptos

- Es un lenguaje creado por Google
- Creado para resolver problemas internos de Google
- Es un lenguaje fuertemente tipado
- Es un lenguaje que hereda su CORE de C++ aunque su sintaxis es mejorada y muy amigable
- Fue pensado para aprovechar los últimos avances de hardware, los multiprocesadores enormes cantidades de memoria y el paralelismo
- Es un lenguaje compilado, generando archivos .EXE portables a cualquier sistema operativo y versión (su EXE contiene todo lo necesario para ejecutarse)
- Obliga a los desarrolladores a llevar buenas prácticas
- Demostró ser el lenguaje ideal para grandes desarrollos con miles y miles de Usuarios concurrentes y millones de transacciones

# Convenciones

- Es fácil de entender, su sintaxis es clara y mejorada con respecto a otros lenguajes
- No hace falta el “;”
- Arroja advertencias, ante variables declaradas no utilizadas y paquetes importados no utilizados
- Las funciones en GO pueden devolver más de un valor
- Se pueden desarrollar en GO tanto instrucciones sincrónicas como asincrónicas.
- Solo existe la instrucción FOR para iteraciones (no existe While, ni Do Until, ni nada similar)
- No es un lenguaje orientado a objetos y resuelve lo que conocemos como POO, con estructuras, funciones, métodos e interfaces
- El scope de las variables, métodos y funciones se determinan con nombres en mayúsculas y minúsculas

# Instalación de GO

Desde el sitio oficial https://golang.org descargar el instalador
Ir al link https://go.dev/doc/install y realizar los pasos en la terminal
Si no tienes instalada una versión anterior entonces no es necesario eliminar go de `/usr/local/go`

# Usando GO

- Para ejecutar con compilador rápido en memoria sin generar un archivo de ejecución del programa, basta con el comando go run main.go
  siempre inicia el programa con package nombre_file

- Para importar librerías se debe usar import “libreria”. Si son más de una librería a importar se debe utilizar:

```
import (
	“libreria1”
	“libreria2”
	“libreria3”
)
```

- Para generar un archivo de ejecución se debe usar el comando:
  go build main.go

# Variables en GO

- Se usa la palabra reservada var para declarar variables se usa

```
var numero int
var texto string
var status bool
```

- Tipos de datos numéricos:
  float32
  float64
  int16
  int8
  int32
  int64
  uint => tipo de entero sin signo
  entre otros...

## Condiciones

Hay 2 formas de sentenciar las condiciones.

- No existen los oarentesis en go. Salvo que hayan que agrupar las condiciones
- Se usan las {} llaves para armar un scope dentro de la condicion

```
if condicion == true {
	... si es true
} else {
	... si es false
}
```

## Mostrar y aceptar datos

Importamos los paquetes:

- "os" => maneja el sistema operativo
- "bufio" =>

```
var numero1 int
var numero2 int

//Con Scanf al dar enter en Linuz es un new line, difiere en Windows

fmt.Println("Ingrese un numero:")
//fmt.Scanf("%d", &numero1)
Scanln(&numero1)

fmt.Println("Ingrese un numero:")
//fmt.Scanf("%d", &numero1)
Scanln(&numero2)

fmt.Printf("%d", numero1+numero2)
```

## Usando Bufio u OS. Alternativa a Scanln

- Le asignamos a la variable scanner un nuevo scanner que va a recibir del standart input (keyboard) un numero

```
...
var resultado int
var leyenda string
...

scanner := bufio.NewScanner(os.Stdin)
if scanner.Scan() {
	leyenda = scanner.Text()
}

resultado = numero1 + numero2
fmt.Println(leyenda, resultado)
```

## Ciclos en Go

Las iteraciones en go solo se realizan por medio del for

- Sintaxis del ciclo for

```
i := 1
for i < 10 {
	fmt.Println(i)
	i++
}
```

o también

```
for i:= 1; i <= 10; i ++ {
	fmt.Println(i)
}
```

- Con el break se puede romper el ciclo for

```
numero := 0
for {
	fmt.Println("Continuo")
	fmt.Println("Ingrese el número secreto")
	fmt.Scanln(&numero)

	if numero == 100 {
		break
	}
}
```

- Se puede usar "continue" sigue en la iteracion sin pasar por las siguientes lineas

```
var i = 0
for i < 10 {
	fmt.Printf("\n Valor de i: %d", i)
	if i == 5 {
		fmt.Printf(" multiplicamos por 2 \n")
		i = i * 2
		continue
	}
	fmt.Printf(" 		Paso por este lugar \n")
	i++
}
```

- El uso Secciones. Se usa "goto" para llamar la sección que es parecido del "continue" solo que va a la sección que se específica

```
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
```

## Funciones

Existen ditintos tipos de funciones. Todo se hace con funciones en Go.

- La estructura de la funcion

```
func uno(numero int) (z int) {
	z = numero * 2
	return
}
```

(Tipando los parametros y lo que la funcion retornará)

- Una función en Go puede retornar 2 valores de diferentes tipos

```
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}
```

- En go el resultado de la funcion, los valores que retorna pueden ser asignados en variables

```
numero, estado := dos(1)
// numero => 5
// estado => true
```

- Las funcionaes pueden recibir un numero indeterminado de parámetros
  En Go. No tiene sobrecarga de funciones o métodos

Con los "...int" le dice que va a recibir un nùmero dinámico o indeterminado de paràmetros enteros

```
func Calculo(numero ...int) int {
	total := 0
	// _ o item key de cada variable dentro de la lista; num valor del elemento en la lista
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}

	return total
}
```

- Range => utiliza una lista o un vector de elementos. Devuelve o retorna 2 valores: el [key] del elemento y el [elemento] que está recorriendo ese momento

## Funciones Anónimas y Closures

- Func Anónimas => son funciones que no tienen nombre. Para crear operaciones sin necesidad de llamarlas en tiempo de ejecución
  (Se crea una variable de tipo función)

```
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma 5 + 7 es = %d \n", Calculo(5, 7))
}
```

- Se realiza esto para luego poder redefinir la variable de tipo funcion

```
Calculo = func(num1 int, num2 int) int {
	return num1 - num2
}

fmt.Printf("Resta 6 - 4 es = %d \n", Calculo(6, 4))
```

- Respetando el numero de paràmetros definido en el tipo de la funcion

- Y cuando una funcion por dentro tiene una función anónima

```
func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13

		return a * b

	}
	fmt.Println(resultado())
}
```

- Closures => Tiene que ver con la protección de código. Es una función que guarda referencias del estado adyacente (àmbito lèxico). En otras palabras, una clausura permite acceder al àmbito de una función exterior desde una función interior. Entonces puede tener un estado único propio. Luego, el estado se aísla a medida que creamos nuevas instancias de la función.

```
func main() {
	TablaDel := 2
	MiTabla := Tabla(TablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
```

## Arreglos estáticos & Slices

## Arreglos

- Se declaran:

```
var tabla [10]int
```

[10]int => dimensión del arrreglo junto con el tipo

- Al declarar sin inicializar valores del arreglo este arreglo sera de [0,0,0,...10 veces]

- Se puede asignar directamente por posiciones
  tabla[0] = 1
  tabla[4] = 15

- Se asigna un arreglo con {}
  tabla := [10]int{5, 6, 7, 8, 12, 2, 3, 5, 11, 12}

- Recorrer un arreglo

```
for i := 0; i < len(tabla); i++ {
	fmt.Println((tabla[i]))
}
```

- Las matrices son un conjunto de vectores-listas-arreglos en 2 dimensiones, filas-columnas

```
var matriz [5][7]int
matriz[3][5] = 1
fmt.Println((matriz))

// resultado [[0 0 0 0 0 0 0] [0 0 0 0 0 0 0] [0 0 0 0 0 0 0] [0 0 0 0 0 1 0] [0 0 0 0 0 0 0]]
```

## Slices

Es similar a una lista o array, que son vectores dinámicos

```
matriz := []int{2, 5, 4}
fmt.Println(matriz)

// resultado [2 5 4]
```

- Pueden ser construidos a partir de un vector normal

```
elementos := [5]int{1, 2, 3, 4, 5}
porcion := elementos[3:]   // desde el tercer elemento hasta el final del vector
porcion2 := elementos[:3]  // desde el primer elemento hasta el tercer elemento del vector
porcion3 := elementos[2:4] // desde el primer elemento hasta el tercer elemento del vector

fmt.Println(porcion)
fmt.Println(porcion2)
fmt.Println(porcion3)

// resultados
// [4 5]
// [1 2 3]
// [3 4]
```

- Utilizando la func make se puede construir slices
  La func make pide 3 parámetros:

1. tipo de slice
2. tamaño del slice
3. capacidad máxima del slice => Esto quiere decir que el slice puede tener un tamaño pero su capacidad máxima es el límite de su tamaño

```
elementos := make([]int, 5, 20) // slice entero de tamaño 5 pero que puede crecer hasta 20
fmt.Println(elementos)

// resultado [0 0 0 0 0]
```

- los Slices pueden ser dimensionado y redimensionados de acuerdo a la necesidad

```
// make => tiene 3 parametros para crear slices: type, len, capacity
elementos := make([]int, 5, 20)
// cap => func que permite revisar la capacidad de un vector
fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

// Resultado Largo 5, Capacidad 20
```

Y luego redimensionado a traves de la func append

```
nums := make([]int, 0, 0)
for i := 0; i < 100; i++ {
	nums = append(nums, i)
}
fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
```

## Mapas

Un Mapa es una estructura que se puede serializar, es decir, que se puede crear en serie
En un mapa se puede crear un par de clave y valor. Que pueden ser string o numerico

- Se declara un mapa de la siguiente manera:

```
// Con la func make
paises := make(map[string]string, 5)

// Sin la func make
campeonato := map[string]int{
	"Barcelona":    39,
	"Real Madrid":  38,
	"Chivas":       37,
	"Boca Juniors": 30,
}
```

- Se agregan elementos:

```
paises["Mexico"] = "D.F."
paises["Argentina"] = "Buenos Aires"

// o en el caso de que la clave sea string
campeonato["River Plate"] = 25
```

- Modificar elementos al map

```
campeonato["Chivas"] = 55
```

- Eliminar elementos del map

```
delete(campeonato, "Real Madrid")
```

- Listando todos los elementos del map

```
for equipo, puntaje := range campeonato {
	fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
}
```

- Buscando elemento

```
puntaje, existe := campeonato["Mineiro"]
fmt.Printf("El puntaje capturado %d, y el equipo existe %t\n", puntaje, existe)

puntaje, existe = campeonato["Chivas"]
fmt.Printf("El puntaje capturado %d, y el equipo existe %t\n", puntaje, existe)
```

## Estructuras (POO en GO)

Es una colección de múltiples campos de datos con sus tipos de datos definidos agrupados. Son útiles para agrupar datos para formar registros personalizados. Consta de ambos incorporado y tipos definidos por el usuario.

- Las estructuras pueden ser modificadas a los largo del programa.

- Ayudan a mejorar la calidad general del código al permitirnos crear y pasar estructuras de datos complejas a través de múltiples módulos. Imagina pasar 10 parámetros a una función, pronto se quedará sin código. En lugar de pasar 10 parátros en una función, ahora simplemente pasa una sola estructura.

- Muy similar a las clases, contiene un conjunto de campos que tienen un tipo definido y un identificador.

**Sintáxis**

```
type usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}
```

## Interfaces

Asi como en la POO, parecidas a las estructuras, pero permiten definir conductas, operaciones y comportamientos. Creando cualquier objeto, ese objeto implementa comportamientos, metodos y funciones.

Las interfaces de Go proporcionan un método para organizar composiciones complejas y aprender a usarlas le permitirá crear código común y reutilizable.

Una interface en Go, es una plantilla de métodos y es utilizado para proporcionar modularidad a Go. Es decir, la interface **indica que métodos deben ser implementados pero no los define**. Las interfaces son útiles ya que definen la especificación de los métodos, de forma que las posibles implementaciones pueden ser intercambiadas.

**Sintáxis**

```
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}
```

## Manejo de archivos desde Go

Para el manejo de archivos podemos usar las siguientes librerías:

- os
- io/ioutil

# Recursión en GO

Las recursiones es un método de resolución de problemas por medio de su división en pequeñas instancias que luego son unidas para generar una solución global

La recursión se comporta como una pila en la que se ingresan valores hasta alcanzar uno o varios casos base a partir de los cuales se solucionará un problema.

Los algoritmos recursivos constan de dos elementos fundamentales:

- Casos base: son los escenarios finales a partir de los cuales se comienzan a agrupar las pequeñas soluciones
- Reglas: las cuales se encargan de dividir el problema principal en diversos casos.

# DEFER - PANIC & RECOVER

- DEFER: es una instrucción que se va a ejecutar siempre que finalice el programa o si una funcion se va a ejecutar y se resuelve por un return o un error

- PANIC: es una instrucción que fuerza a un error y el sistema aborta la ejecución

# GO RUTINES

Es un hilo o thread de ejecución ligero. supongamos que tenemos una llamada a la funcion ```f(s)``` Así es como la llamaríamos de la manera tradicional, o de manera síncrona. 

Para llamar esta función en una **gorutine**, usamos ``` go f(s)```. Esta nueva gorutine se ejecutará de manera concurrente junto con otras funciones o métodos. 

Incluso, la función principal main se ejecuta dentro de una **gorutine**.

# CHANNELS (Diálogo entre Gorutines)

Los canales es un espacio de memoria, de dialogo en el que se van a comunicar las **gorutines** concurrentes. Puedes enviar valores por un canal de una gorutine y recibir esos valores en otra gorutine.