// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	fmt.Println("Enter first num")
	fmt.Scan(&num1)
	fmt.Println("Enter second num")
	fmt.Scan(&num2)

	var operator string
	fmt.Println("Choose an operator. + - x /")
	fmt.Scan(&operator)

	if operator == "+" {
		result := num1 + num2
		fmt.Println(result)
	} else if operator == "-" {
		result := num1 - num2
		fmt.Println(result)
	} else if operator == "x" {
		result := num1 * num2
		fmt.Println(result)
	} else if operator == "/" {
		if num2 == 0 {
			fmt.Println("Invalid!! cannot divide by 0")
		} else {
			result := num1 / num2
			fmt.Println(result)
		}
	} else {
		fmt.Println("not an valid operator")
	}
}
