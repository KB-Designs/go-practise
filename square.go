package main
import "fmt"

func square(num int)int{
	squared:=num*num
	return squared
}

func main (){
	var num2besqured int
	fmt.Print("Enter a number\n")
	fmt.Scan(&num2besqured)

	x:=square(num2besqured)
	fmt.Println(x)
}