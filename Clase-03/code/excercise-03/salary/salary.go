package salary

func CalculateSalary(minutesWorked int, category string) (salaryTotal float64) {
	hourlyRate := 0.0
	baseSalary := float64(minutesWorked) / 60.0

	switch category {
	case "C":
		hourlyRate = 1000.0
	case "B":
		hourlyRate = 1500.0
		baseSalary *= 1.2
	case "A":
		hourlyRate = 3000.0
		baseSalary *= 1.5
	}

	salaryTotal = baseSalary * hourlyRate

	return
}
