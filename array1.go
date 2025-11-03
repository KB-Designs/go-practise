package main
import "fmt"

func main (){
	nums:=[...]int {5,4,3,2,1}

	slice:=nums[0:3]
	fmt.Println(slice)

	// empty slice

	


	//slice
	


	slice1:=[]int{33,22,11,44,55}
	slice1=append(slice1,99)

	for index,value:= range slice1{
		fmt.Printf("index: %v value:%v\n",index,value)
	}
	

	fmt.Println(slice1)

}