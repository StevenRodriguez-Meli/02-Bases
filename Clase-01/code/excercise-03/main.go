package main

import "fmt"

var firstName string
var lastName string
var age int

var driverLicense = true
var person, height int

func main() {

	lastName2 := 6
	childsNumber := 2

	fmt.Println("La primera variable es incorrecta ya que las variables en go no pueden comenzar con un numero:", firstName)
	fmt.Println("La segunda variable esta inicializada correctamente:", lastName)
	fmt.Println("La tercera variable es incorrecta ya que la palabra clave int no debe preceder a age en la declaración de la variable:", age)
	fmt.Println("La cuarta variable tambien es incorrecta las variables no pueden comenzar con numero:", lastName2)
	fmt.Println("La quinta variable se declara correctamente aunque se podria mejorar usando camelCase dejando la variable como driverLicense:", driverLicense)
	fmt.Println("La sexta variable es incorrecta, si queremos declarar dos variables en la misma linea se necesita separar por ,: ", person, height)
	fmt.Println("La septima variable es correcta: ", childsNumber)
}
