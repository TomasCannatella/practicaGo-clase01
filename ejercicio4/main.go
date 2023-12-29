/*
	Ejercicio 4 - Tipos de datos

	Un estudiante de programación intentó realizar declaraciones de variables
	con sus respectivos tipos en Go, pero tuvo varias dudas mientras lo hacía.
	A partir de esto, nos brindó su código y pidió la ayuda de un desarrollador
	experimentado que pueda:
		- Verificar su código y realizar las correcciones necesarias.
*/
var firstName string = "Mary" //Esta bien

var lastName string = "Smith" // Esta bien

var age int = "35" //Esta bien

boolean := "false"
/*
   Esta mal, porque para indicar un valor valor booleano
   se hace sin el uso de las comillas, ya que en este caso
   go lo interpretaria como un dato del tipo string, ademas
   boolean no es una palabra reservada de go ya que en
   go para indicar que una variable es del tipo booleana se
   utiliza la palabra reservada bool,podría quedar asi:
   hasMarried := false
*/
var salary string = 45857.90
/*
	Esta mal ya que el valor que se le esta asignando a la variable es un
	valor del tipo float, entonces la declaracion de la misma quedaria asi:
	var salary float = 45857.90
*/