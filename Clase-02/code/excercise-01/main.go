package main

import (
	"fmt"
)

var word string

func main() {

	fmt.Print("Enter a word: ")
	fmt.Scan(&word)

	letterCount := 0
	for i := 0; i < len(word); i++ {
		letterCount++
	}
	fmt.Printf("The word \"%s\" has %d letters.\n", word, letterCount)

	fmt.Println("Letters:")
	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))
	}
}
