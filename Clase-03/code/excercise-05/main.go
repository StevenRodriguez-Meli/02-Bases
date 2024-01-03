package main

import (
	"app/Clase-03/code/excercise-05/animals"
	"fmt"
)

const (
	dog = "dog"
	cat = "cat"
)

func main() {
	animalDog, msg := animals.Animal(dog)
	if msg != "" {
		fmt.Println("Error:", msg)
		return
	}

	animalCat, msg := animals.Animal("pig")
	if msg != "" {
		fmt.Println("Error:", msg)
		return
	}

	var amount float64

	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Println("Total de alimento necesario:", amount, "kg")
}
