package main

import (
	"errors"
	"fmt"
)

var ErrSalary = errors.New("error: salary is less than 10000")

func checkSalaryError(salary int) error {
	if salary <= 10000 {
		return ErrSalary
	}
	return nil
}
func main() {
	salary := 90000

	err := checkSalaryError(salary)
	if errors.Is(err, ErrSalary) {
		fmt.Println(err)
	} else {
		fmt.Println("Salary is acceptable")
	}
}
