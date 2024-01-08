package main

import (
	"app/Clase-07/code/excercise-05/salary"
	"fmt"
)

func main() {
	hoursWorked := -100
	hourlyRate := 1500

	salary, err := salary.CalculateSalary(hoursWorked, hourlyRate)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Monthly salary: $%.2f\n", salary)
	}
}
