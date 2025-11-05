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