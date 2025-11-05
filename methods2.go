package main
import "fmt"

type rectangle struct{
	length float64
	width float64
}

func (r rectangle) area()float64{
	area:=r.length*r.width
	return area
}

func main (){
	rec1:=rectangle{22 ,10}
	area1:=rec1.area()
	fmt.Println("Area is:", area1)
}