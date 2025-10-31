package main
import "fmt"

//Example 1 — Assign to a Variable




func main (){
	sum := func(x, y int) int {
    return x + y
	}

	fmt.Println(sum(3,4))
} 
	


//Example 1 — Assign to a Variable


package main
import "fmt"

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
    firstResult := arithmetic(a, b)
    secondResult := arithmetic(firstResult, c)
    return secondResult
}

func main() {
    // Anonymous function for addition
    fmt.Println(aggregate(2, 3, 4, func(x, y int) int {
        return x + y
    })) 

    
    fmt.Println(aggregate(2, 3, 4, func(x, y int) int {
        return x * y
    })) 
}

