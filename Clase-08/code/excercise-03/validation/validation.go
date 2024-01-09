package validation

import (
	"app/Clase-08/code/excercise-03/client"
	"errors"
)

func CheckClientExists(name string) error {
	_, exists := client.ClientDatabase[name]
	if exists {
		return errors.New("client already exists")
	}
	return nil
}

func ValidateClientData(client client.Client) error {
	if client.ID == 0 || client.Name == "" || client.PhoneNumber == "" || client.Home == "" {
		return errors.New("all fields must be filled")
	}
	return nil
}
