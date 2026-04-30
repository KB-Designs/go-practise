package main

import (
	"fmt"
)

func main() {
	var word string
	reversed := ""

	fmt.Println("Enter a word: ")
	fmt.Scanln(&word)

	for i := len(word) - 1; i >= 0; i-- {
		reversed = reversed + string(word[i])
	}
	fmt.Println("Reversed word: ", reversed)

	if word == reversed {
		fmt.Println("The word is a palindrome.")
	} else {
		fmt.Println("The word is not a palindrome.")
	}
}
