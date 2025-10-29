package main
import "fmt"

func squareperi(num int)(int,int){
	squared:=num*num
	perimeter:=num*4
	return squared,perimeter
	
	
}

func main (){
	var num2besqured int
	fmt.Print("Enter a number\n")
	fmt.Scan(&num2besqured)

	squaredsquare,perimetersquare:= squareperi(num2besqured)

	fmt.Printf("square : %v\nperimeter : %v\n",squaredsquare,perimetersquare)
}