/*
Ejercicio 2 - Clima

Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura,
																	   humedad y
																	   presión atmosférica de distintos lugares.

Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
Imprimí los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?
	- En este caso para la temperatura utiliza un tipo de dato Integer Signed de 8bits ya que podemos tener temperaturas bajo cero pero ninguna esta por encima del valor 127 o -128
	- Para la humedad asigne un tipo de dato Integer unsigned ya que siempre va a ser un valor positivo que va desde el 0 hasta el 100
	- Y para la presion atmosferica le asigne un tipo de dato Integer de 16bits ya que los valores siempre van a ser positivo y por encima de 128.
*/

package main

import "fmt"

func main() {
	var temperatura int8 = 29
	var humedad uint8 = 42
	var presionAtmosferica uint16 = 1011
	fmt.Print("En Mendoza hay una temperatura de ", temperatura, "° ")
	fmt.Print("humedad: ", humedad, "%")
	fmt.Println(" y una presion atmosférica: ", presionAtmosferica, "hpa")
}
