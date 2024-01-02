package main

import "fmt"

var monthNumber int

func main() {
	months := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	fmt.Print("Enter the month number: ")
	fmt.Scan(&monthNumber)

	if monthNumber >= 1 && monthNumber <= 12 {

		fmt.Printf("%d, %s\n", monthNumber, months[monthNumber-1])

	} else {
		fmt.Println("Invalid month number")
	}

	fmt.Println("This way seems the most appropriate to me since it simplifies the problem into fewer lines of code.")
}
