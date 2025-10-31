package main
import (
	"math"
  "fmt"
)

// Define specialComputation() herex float64
func specialComputation(x float64) float64{
  return math.Log2(math.Sqrt(math.Tan(x)))

}

func main() {
  var a, b, c, d float64
  a = .0214
  b = 1.02
  c = 0.312
  d = 4.001
  
  // Replace the following four lines with specialComputation()
 a=specialComputation(a)
 b=specialComputation(b)
 c=specialComputation(c)
 d=specialComputation(d)

  fmt.Println(a, b, c, d)
}


func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
  averageGrade := (midtermGrade + finalGrade) / 2
  var gradeLetter string
  if averageGrade > 90 {
    gradeLetter = "A"
  } else if averageGrade > 80 {
    gradeLetter = "B"
  } else if averageGrade > 70 {
    gradeLetter = "C"
  } else if averageGrade > 60 {
    gradeLetter = "D"
  } else {
    gradeLetter = "F"
  }

  return gradeLetter, averageGrade 
}

func main() {
  var myMidterm, myFinal float32
  myMidterm = 89.4
  myFinal = 74.9
  var myAverage float32
  var myGrade string
  myGrade, myAverage = GPA(myMidterm, myFinal)
  fmt.Println(myAverage, myGrade) // Prints 82.12 B
}

