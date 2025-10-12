package main

import (
	"fmt"
)

func main() {
	array1 := [...]int{12, 3, 4, 5}

	fmt.Println(array1[2])
	fmt.Println(len(array1))

	var name= "Brian"
	name2 := "kimani"

	fmt.Println(name)
	fmt.Println(name2)

	//change array element
	array1[1]=50

	fmt.Println(array1)
}
