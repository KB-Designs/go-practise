package main 
import "fmt"

func main (){
	grades:=[]int{65,76,66,88}

	sum:=0
	for _,v:=range grades{
		
		sum+=v

		
	}
	fmt.Println(sum)
	
}


Write a short Go program that:

Defines a Person struct with name and age

Creates a method birthday() that increases age by 1 using a pointer receiver

Prints the new age after calling the method