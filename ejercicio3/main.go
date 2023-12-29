var 1firstName string 
/*
	Esta mal porque el nombre de la variable comienza con un numero entonces la misma
	se podria nombrar de esta forma firstName y ademas en este ejemplo no tiene
	asignado ningun valor. En GO no pasa nada porque de forma automatica lo dectecta 
	como un string vacio pero estaria bueno ponerle algún nombre,
	entonces la declarcion quedaria asi: 
	var firstName string := "Jonh"
*/
var lastName string 
/*
	Esta bien pero  no tiene asignado ningun valor. En GO no pasa
	nada porque de forma automatica lo dectecta como un string vacio pero
	estaria bueno ponerle algún apellido
	var lastName string := "Doe"
*/

var int age // var age int := 21

1lastName := 6 ?
/*
Esta mal porque el nombre de la variable comienza con un numero, ademas
el nombre o el valor que le puso a la misma no seria el correcto porque
a ver el codigo a simple vista no se entiende que se esta tratando de 
declarar. A mi pensar interpreto que capaz quiere declarar la cantidad
de cuentas que tiene en un banco y se podria de clarar de dos formas
de manera explicita y implicita, quedaria algo asi:
	var countBankAccount uint8 = 6
	countBankAccount := 6
*/

var driver_license = true
/* 
	Esta bien pero se puede mejorar
	en este caso para que quede mas prolijo y como buena practica se podria utilizar
	camelcase y ademas declararla de manera explicita, quedaria asi:
	var hasDriverLicense boolean = true
	o en caso de querer declararla de forma implicita directamente podria quedar asi:
	hasDriverLicense boolean := true
*/

var person height int
/*
	Esta mal, ya que una variable no puede contener espacio
	se podria declarar de esta forma
	var pesonHeight float := 55.0
*/

childsNumber := 2 
/*
	Esta bien declarada ya que la declaro de manera implicita, 
	go es capaz de distinguir los tipos de datos.
*/
