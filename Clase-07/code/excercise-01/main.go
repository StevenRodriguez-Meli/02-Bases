package main

import (
	"fmt"
)

type SalaryError struct {
	message string
}

func (e SalaryError) Error() string {
	return e.message
}

func main() {
	salary := 180000

	if salary < 150000 {
		err := SalaryError{"Error: the salary entered does not reach the taxable minimum"}
		fmt.Println(err)
	} else {
		fmt.Println("Must pay tax")
	}
}
