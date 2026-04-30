// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	countdown := 3
	for countdown > 0 {
		fmt.Printf("countdown:%d\n", countdown)
		countdown--
	}

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Println(idx, val)
	}

	var word string
	fmt.Println("Enter your name: ")
	fmt.Scan(&word)

	for i := len(word) - 1; i >= 0; i-- {
		fmt.Println(string(word[i]))

	}

	str := "brian"
	reversed := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	fmt.Println(reversed)

}
