package main 
import "fmt"

func radius()(int int){
	var x int=2
	var y int=4
	return x,y
}

func main (){
	r1,r2:=radius()

	fmt.Println(r1)
	fmt.Println(r2)

}