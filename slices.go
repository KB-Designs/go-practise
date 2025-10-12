package main
import ("fmt")

func main (){
	//using the []datatypes {values}
	myslice :=[]int{23,4,5,3}

	//from an array
	myarray:=[...]int {3,44,55,32,23}
	myslice1 := myarray[0:3]

	fmt.Println(myslice)
	fmt.Println(myslice1)
}