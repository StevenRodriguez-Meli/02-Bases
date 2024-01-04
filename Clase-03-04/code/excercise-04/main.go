package main

import (
	"app/Clase-03-04/code/excercise-04/calculations"
	"fmt"
)

func main() {
	minFunc, err := calculations.Operation(calculations.Minimum)
	if err != "" {
		fmt.Println("Error:", err)
		return
	}

	averageFunc, err := calculations.Operation(calculations.Average)
	if err != "" {
		fmt.Println("Error:", err)
		return
	}

	maxFunc, err := calculations.Operation(calculations.Maximum)
	if err != "" {
		fmt.Println("Error:", err)
		return
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Minimum:", minValue)
	fmt.Println("Average:", averageValue)
	fmt.Println("Maximum:", maxValue)
}
