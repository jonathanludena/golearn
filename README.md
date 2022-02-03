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