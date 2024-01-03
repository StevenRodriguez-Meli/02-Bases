package main

import (
	"app/Clase-03/code/excercise-02/grades"
	"fmt"
)

func main() {
	gradesList := []int{80, 75, 90, 100, -5, 85}
	average := grades.CalculateAverage(gradesList...)

	fmt.Printf("The average is: %.2f\n", average)
}
