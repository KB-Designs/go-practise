package main
import "fmt"

// arithmetic function for addition
func add(x, y int) int {
    return x + y
}

// arithmetic function for multiplication
func multiply(x, y int) int {
    return x * y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
    firstResult := arithmetic(a, b)
    secondResult := arithmetic(firstResult, c)
    return secondResult
}

func main() {
    fmt.Println(aggregate(2, 3, 4, add))      // (2 + 3) + 4 = 9
    fmt.Println(aggregate(2, 3, 4, multiply)) // (2 * 3) * 4 = 24
}
