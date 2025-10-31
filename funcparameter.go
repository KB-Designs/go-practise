package main
import "fmt"

func add(x,y int)int{
	return x+y
}

func multiply(x,y int)int{
	return x*y
}

func maths(a,b,c int,arith func(int,int)int) int{
	firstresult:=a+b
	secondresult:=firstresult+c

	return secondresult
}

func main (){
	answer1:=maths(2,4,6,add)
	answer2:=maths(4,6,5,multiply)

	fmt.Println(answer1)
	fmt.Println(answer2)
}