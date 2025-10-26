package main 
import "fmt"

func double(num int)int{
	return num*2
}
func startGame() {
  instructions := "Press enter to start..." 
  fmt.Println(instructions) 
}

func main (){
	x:=double(3)
	fmt.Println(x)
	fmt.Println(double(4))

	startGame()
}