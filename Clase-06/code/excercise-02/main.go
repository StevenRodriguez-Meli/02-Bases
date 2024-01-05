package main

import (
	"app/Clase-06/code/excercise-02/product"
	"fmt"
)

func main() {
	product1, err1 := product.CreateProduct("Small", 2000)
	if err1 != "" {
		fmt.Println("Error:", err1)
		return
	}
	product2, err2 := product.CreateProduct("Medium", 3400)
	if err2 != "" {
		fmt.Println("Error:", err2)
		return
	}
	product3, err3 := product.CreateProduct("Large", 5600)
	if err3 != "" {
		fmt.Println("Error:", err3)
		return
	}

	fmt.Println("Price of Small:", product1.Price())
	fmt.Println("Price of Medium:", product2.Price())
	fmt.Println("Price of Large:", product3.Price())
}
