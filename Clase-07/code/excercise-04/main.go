package main

import (
	"errors"
	"fmt"
)

var ErrSalary = errors.New("error: salary is less than 10000")
var ErrMinTaxableAmount error

func checkSalaryError(salary int) error {
	if salary <= 10000 {
		return ErrSalary
	} else if salary < 150000 {
		ErrMinTaxableAmount = fmt.Errorf("error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
		return ErrMinTaxableAmount
	}
	return nil
}

func main() {
	salary := 900000

	err := checkSalaryError(salary)
	if err != nil {
		if errors.Is(err, ErrSalary) {
			fmt.Println(err)
		} else if errors.Is(err, ErrMinTaxableAmount) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Salary is acceptable")
	}
}
