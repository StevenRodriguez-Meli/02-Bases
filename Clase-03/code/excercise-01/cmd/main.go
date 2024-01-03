package main

import (
	"app/Clase-03/code/excercise-01/taxes"
	"fmt"
)

func main() {
	salary := 60000.0
	tax := taxes.CalculateTax(salary)

	fmt.Printf("The tax to be paid for a salary of $%.2f is: $%.2f\n", salary, tax)
}
