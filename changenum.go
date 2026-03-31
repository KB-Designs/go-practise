package main

import "fmt"

func main() {
	num := 10
	fmt.Println(num)

	changenum(&num)

	fmt.Println(num)

}

func changenum(x *int) {
	*x = 20
}
