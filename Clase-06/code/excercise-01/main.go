package main

import (
	"app/Clase-06/code/excercise-01/student"
)

func main() {
	student1 := &student.Student{
		Name:     "John",
		LastName: "Doe",
		DNI:      "12345678",
		Date:     "01-01-1990",
	}

	student1.Details()
}
