package main
import "fmt"

func sum(num1 int, num2 int) (result int) {
	result = num1+num2
	return result
	
}

func main (){
	y:=sum(2,3)
	fmt.Printf("sum is :%v" ,y)
}