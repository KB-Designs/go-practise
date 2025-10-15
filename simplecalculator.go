package main 
import "fmt"

func simplecalculator(){
	fmt.Println("welcome simple calculator")
	
	var num1,num2 float64
	

	fmt.Print("enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("enter second number: ")
	fmt.Scan(&num2)

	var operator string
	fmt.Print("choose an operator among +,-,*,/:\n")
	fmt.Scan(&operator)
	
	if operator=="+"{
		result:=num1+num2
		fmt.Printf("answer :%v\n",result)
	}else if operator=="-"{
		result:=num1-num2
		fmt.Printf("answer :%v\n",result)
	}else if operator=="*"{
		result:=num1*num2
		fmt.Printf("answer :%v\n",result)
	}else if operator=="/"{
		if num2!=0{
			result:=num1/num2
		    fmt.Printf("answer :%v\n", result)
		} else {
			fmt.Printf("Error: Division by zero is not allowed\n")
		}	
	}else{
		fmt.Printf("Invalid operator\n")
	}
}

func main (){
		simplecalculator()
	}
	