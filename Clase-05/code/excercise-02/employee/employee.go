package employee

import (
	"fmt"
	"time"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Position:", e.Position)
	fmt.Println("Person ID:", e.Person.ID)
	fmt.Println("Name:", e.Person.Name)
	fmt.Println("Date of Birth:", e.Person.DateOfBirth.Format("02-01-2006"))
}
