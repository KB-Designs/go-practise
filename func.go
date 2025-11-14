package main

import "fmt"

type person struct {
	name string
	age  int
	city string
}

func main() {
	p1 := person{
		name: "brian",
		age:  22,
		city: "thika",
	}

	fmt.Println(p1)
}
