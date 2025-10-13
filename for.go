package main
import ("fmt")

func main (){
	

	for t:=1; t<3; t++{
		p:=2000
		r:=0.12
		

		interest:=float64(p)*r*float64(t)

		fmt.Println(t)
		fmt.Println(interest)
	}

}