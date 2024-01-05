package product

import "fmt"

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{Id: 1, Name: "Product 1", Price: 19.99, Description: "Description of Product 1", Category: "Category 1"},
	{Id: 2, Name: "Product 2", Price: 29.99, Description: "Description of Product 2", Category: "Category 2"},
}

func Save(newProduct Product) {
	Products = append(Products, newProduct)
}

func GetAll() {
	fmt.Println("List of Products:")
	for _, p := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", p.Id, p.Name, p.Price, p.Description, p.Category)
	}
}
