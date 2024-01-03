package main

import (
	"app/Clase-03/code/excercise-03/salary"
	"fmt"
)

func main() {
	minutes := 160 * 60
	category := "A"

	salaryAmount := salary.CalculateSalary(minutes, category)

	fmt.Printf("The salary for Category %s is: $%.2f\n", category, salaryAmount)
}
