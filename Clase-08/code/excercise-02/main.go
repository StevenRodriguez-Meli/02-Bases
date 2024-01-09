package main

import (
	"app/Clase-08/code/excercise-02/createFile"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "customers.txt"

	createFile.CreateFileWithCustomerData(fileName)

	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("The indicated file '%s' was not found or is damaged", fileName))
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("Error while reading file: %s", err))
	}

	fmt.Printf("Contenido del archivo '%s':\n", fileName)
	fmt.Println(string(bytes))

	fmt.Println("Ejecución finalizada")
}
