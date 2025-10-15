package main 
import "fmt"

func recursion(num int) (result int){
	if num==0 || num ==1 {
		result=1
		return result
	} else {
		result= num * recursion(num-1)
		return result
	}
}

func main () {
	x:=recursion(10)

	fmt.Println(x)
}
