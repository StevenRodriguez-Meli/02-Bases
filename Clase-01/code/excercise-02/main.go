package main

import "fmt"

var temperature float64
var pressure int
var humidity float64

func main() {

	temperature = 25.5
	pressure = 1000
	humidity = 50.0

	fmt.Printf("Temperature: %.2f°C\n", temperature)
	fmt.Printf("Pressure: %d mbar\n", pressure)
	fmt.Printf("Humidity: %.2f%%\n", humidity)
}
