package main 
import "fmt"

type person struct{
	name string 
	age int
	city string
}

type personlist []person


func (ptr *personlist) add(name,city string ,age int){
	newperson:=person{
		name:name,
		city:city,
		age:age,
		
	}
	*ptr=append(*ptr, newperson)
}

func main(){
	var person1 personlist

	person1.add("kimani","ruiru",22)

	fmt.Println(person1)
}
