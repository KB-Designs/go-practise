package main 
import "fmt"


type person struct {
	name string
	age int
	city string
	grades []int
}

func main (){
	p1:=person {"kimani",22,"nairobi",[]int{79}}
	p1.age=23

	p3:=person{
		name:"kimani",
		age:40,
		city:"malaa",
		grades:[]int{2255,66,77},
	}

	fmt.Println(p1)
	fmt.Println(p3)
}