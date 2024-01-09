package createFile

import (
	"fmt"
	"os"
)

func CreateFileWithCustomerData(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error creating file '%s'", fileName))
	}
	defer file.Close()

	customerData := `
Nombre: Juan Lozano
Edad: 30
Email: juan@example.com

Nombre: María Gutierrez
Edad: 25
Email: maria@example.com
	`

	_, err = file.WriteString(customerData)
	if err != nil {
		panic(fmt.Sprintf("Error writing to file: %s", err))
	}

	fmt.Printf("Archivo '%s' creado con información de clientes.\n", fileName)
}
