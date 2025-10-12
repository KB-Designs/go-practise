package main

import (
	"fmt"
)

func main() {
	var num2 int = 20
	var name string = "Brian"
	num1 := 5

	fmt.Println(num2)
	fmt.Println(name)
	fmt.Println(num1)

	var array1 = [4]int{3, 4, 5, 6}
	array2 := [...]int{2, 5, 6, 7, 7}

	fmt.Println(array1)
	fmt.Println(array2)

}
