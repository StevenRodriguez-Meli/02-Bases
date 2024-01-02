package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {

	fmt.Println("Age of Benjamin:", employees["Benjamin"])

	over21 := 0
	for _, age := range employees {
		if age > 21 {
			over21++
		}
	}
	fmt.Println("Employees older than 21:", over21)

	employees["Federico"] = 25
	fmt.Println("Employee Federico added successfully.")

	delete(employees, "Pedro")
	fmt.Println("Employee Pedro removed from the records.")

	fmt.Println("Updated list of employees and ages:")
	for name, age := range employees {
		fmt.Printf("%s: %d years\n", name, age)
	}
}
