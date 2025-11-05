package main
import "fmt"





func main (){
	slice:=[]int{33,44,100,66,77}

	max:=0

	for _,value:= range slice{
		if max<value{
			max=value
			
		}
    
	}
	fmt.Println("The maximum value is:", max)
}