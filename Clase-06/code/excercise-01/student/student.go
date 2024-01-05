package student

import (
	"fmt"
)

type Student struct {
	Name     string
	LastName string
	DNI      string
	Date     string
}

func (student *Student) Details() {
	fmt.Println("Name:", student.Name)
	fmt.Println("LastName:", student.LastName)
	fmt.Println("DNI:", student.DNI)
	fmt.Println("Date:", student.Date)
}
