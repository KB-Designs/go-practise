package main
import "fmt"

func sum(num1 int, num2 int) (result int) {
	result = num1+num2
	return result
	
}

func substraction(num1 int, num2 int) (result int) {
	result = num1+num2
	return result
}

func multiplication(num1 int ,num2 int) (result int) {
	result = num1+num2
	return result
}


func main (){
	y:=sum(2,3)
	x:=substraction(3,4)
	z:=multiplication(4,6)


	fmt.Printf("sum is :%v\n" ,y)
	fmt.Printf("sum is :%v\n" ,x)
	fmt.Printf("sum is :%v" ,z)

}