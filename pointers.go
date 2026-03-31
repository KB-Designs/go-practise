package main

import "fmt"

func main() {
	age := 18
	ageptr := &age

	fmt.Println(age)
	fmt.Println(ageptr)

	*ageptr = 30
	fmt.Println(age)
}
