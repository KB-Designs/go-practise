package main

import (
	"fmt"
)

func main() {

	//for t := 1; t < 3; t++ {
	//p:=2000
	//r:=0.12

	//interest:=float64(p)*r*float64(t)

	//fmt.Println(t)
	//fmt.Println(interest)

	var name string
	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	for i := 0; i < len(name); i++ {
		//fmt.Println(len(name))
		fmt.Println(string(name[i]))
		//break

	}
	//fmt.Println(string(name[5]))
}
