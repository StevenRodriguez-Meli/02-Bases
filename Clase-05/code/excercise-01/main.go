package main

import (
	"app/Clase-05/code/excercise-01/product"
	"fmt"
)

func main() {

	newProduct := product.Product{
		Id:          3,
		Name:        "New Product",
		Price:       39.99,
		Description: "Description of New Product",
		Category:    "Category 3",
	}

	product.Save(newProduct)

	product.GetAll()

	productID := 4

	foundProduct := product.GetByID(productID)
	if foundProduct != (product.Product{}) {
		fmt.Printf("\nProducto found with ID %d:\nName: %s, Price: %.2f, Description: %s, Category: %s\n", foundProduct.Id, foundProduct.Name, foundProduct.Price, foundProduct.Description, foundProduct.Category)
	} else {
		fmt.Printf("\nProduct not found with ID%d\n", productID)
	}
}
