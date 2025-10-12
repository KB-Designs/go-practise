package main
import ("fmt")

func main (){
	var name string= "Brian"
	var age int = 22

	color :="white"

	fmt.Printf("I am %v, just turned %v\n", name,age)
	fmt.Printf("name :%v, type: %T\n",name, name)
	fmt.Printf("age :%v, type:%T\n", age,age)
	fmt.Printf("color :%v, type:%T", color,color)
}