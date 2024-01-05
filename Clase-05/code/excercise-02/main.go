package main

import (
	"app/Clase-05/code/excercise-02/employee"
	"time"
)

func main() {
	person := employee.Person{
		ID:          1,
		Name:        "John Doe",
		DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	employee := employee.Employee{
		ID:       1001,
		Position: "Manager",
		Person:   person,
	}

	employee.PrintEmployee()
}
