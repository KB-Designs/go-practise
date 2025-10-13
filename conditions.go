package main
import ("fmt")

func main (){
	/*
	age:=12
	
	if age<18 {
		fmt.Println("not eligible to vote")
	} else if age>= 18{
		fmt.Println("eligible to vote")
	}else if age>120 {
		fmt.Println("vote from home")
	}else {
		fmt.Println("Not eligiblr to vote")
	}
		*/

	marks:= 0
	
	if marks>=70 && marks <=100 {
		fmt.Println("Grade A")
	} else if marks >=60 && marks<=69 {
		fmt.Println("Grade B")
	} else if marks>=50 && marks<=59{
		fmt.Println("Grade C") 
	} else if marks >= 40 && marks <=49 {
		fmt.Println("Grade D")
	}else if marks <40 && marks >0 {
		fmt.Println("Grade E")
	} else {
		fmt.Println("Enter marks between 1 and 100")
	}

	myarray :=[...] int {44,56,77,56,76}
	myslice1 := []int {1,2,3,4,5}
	myslice2 := myarray[1:3]



	
	fmt.Println(myslice1)
	fmt.Println(myslice2)
}
