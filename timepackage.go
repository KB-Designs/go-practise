package main
import (
	"fmt"
	"time"
)

func main (){
	var yearbirth int
	fmt.Print("Year of birth:\n")
	fmt.Scan(&yearbirth)

	currenttime:= time.Now()
	currentyear:=currenttime.Year()

	years:=currentyear-yearbirth

	fmt.Printf("On %v your are %v's old\n",currenttime,years)
}