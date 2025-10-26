package main
import "fmt"

// computeMarsYears takes earthYears as a parameter
func computeMarsYears(earthYears int) int {
    earthDays := earthYears * 365
    marsYears := earthDays / 687
    return marsYears
}

func main() {
    myAge := 25

    // Call computeMarsYears with myAge
    myMartianAge := computeMarsYears(myAge)
    fmt.Println(myMartianAge)
}
