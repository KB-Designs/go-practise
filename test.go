package main 
import "fmt"

func test()(x int,y int){
	return
}

func main (){
	a,b:=test()

	fmt.Println(a)
	fmt.Println(b)
}