package client

type Client struct {
	File        string
	Name        string
	ID          int
	PhoneNumber string
	Home        string
}

var ClientDatabase = make(map[string]Client)
