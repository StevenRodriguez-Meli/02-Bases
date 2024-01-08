package salary

import "errors"

const minHoursPerMonth = 80
const minTaxableAmount = 150000
const taxRate = 0.1

func CalculateSalary(hoursWorked, hourlyRate int) (float64, error) {
	if hoursWorked < minHoursPerMonth || hoursWorked < 0 {
		return 0, errors.New("error: the worker cannot have worked less than 80 hours per month")
	}

	salary := float64(hoursWorked * hourlyRate)

	if salary >= minTaxableAmount {
		salary -= salary * taxRate
	}

	return salary, nil
}
