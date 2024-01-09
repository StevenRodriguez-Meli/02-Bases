package main

import (
	"app/Clase-08/code/excercise-03/client"
	"app/Clase-08/code/excercise-03/validation"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("End of execution")
			fmt.Println("Several errors were detected at runtime")
		} else {
			fmt.Println("End of execution")
		}
	}()

	existingClient := client.Client{
		File:        "file",
		Name:        "John Doe",
		ID:          12345,
		PhoneNumber: "123456789",
		Home:        "123 Main St",
	}
	client.ClientDatabase[existingClient.Name] = existingClient

	newClient := client.Client{
		File:        "file",
		Name:        "hola",
		ID:          3232,
		PhoneNumber: "123456789",
		Home:        "123 Main St",
	}

	err := validation.CheckClientExists(newClient.Name)
	if err != nil {
		panic("Error: client already exists")
	}

	err = validation.ValidateClientData(newClient)
	if err != nil {
		panic(err)
	}

	client.ClientDatabase[newClient.Name] = newClient
	fmt.Println("Client registered successfully")
}
