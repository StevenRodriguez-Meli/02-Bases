package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "customer.txt"

	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("The indicated file '%s' was not found or is damaged", fileName))
	}
	defer file.Close()

	fmt.Println("Archivo leído exitosamente.")
	fmt.Println("Ejecución finalizada")
}
