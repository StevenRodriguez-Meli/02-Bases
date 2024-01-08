package main

import (
	"fmt"
)

type SalaryError struct {
	message string
}

func (e *SalaryError) Error() string {
	return e.message
}

func Is(err error) bool {
	_, ok := err.(*SalaryError)
	return ok
}

func main() {
	salary := 90000

	if salary <= 10000 {
		err := &SalaryError{"Error: salary is less than 10000"}
		if Is(err) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Salary is acceptable")
	}
}
