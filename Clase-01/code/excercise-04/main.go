package main

import "fmt"

var lastName string = "Smith"
var age int = 35
var salary string = "45857.90"
var firstName string = "Mary"

func main() {

	boolean := false

	fmt.Println("La primera variable esta declarada correctamente", lastName)
	fmt.Println("La segunda variable esta declarada incorrectamente ya que dice que es un entero y se le asigna un string:", age)
	fmt.Println("La cuarta variable es incorrecta ya que es un string y se le asigna un int:", salary)
	fmt.Println("La tercera variable es incorrecta ya que boolean es de tipo booleano y se le asigna un string:", boolean)
	fmt.Println("La quinta variable se declara correctamente:", firstName)
}
