package main

import "fmt"

var age, yearsAtJob int
var employed bool
var salary float64

func main() {

	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	fmt.Print("Are you employed? (true/false): ")
	fmt.Scan(&employed)

	fmt.Print("Enter years at current job: ")
	fmt.Scan(&yearsAtJob)

	fmt.Print("Enter your salary without points: ")
	fmt.Scan(&salary)

	if age > 22 && employed && yearsAtJob > 1 {
		if salary > 100000 {
			fmt.Println("Congratulations! You qualify for an interest-free loan.")
		} else {
			fmt.Println("You qualify for a loan, but interest will be charged.")
		}
	} else {
		fmt.Println("Sorry, you don't meet the criteria for a loan.")
	}
}
