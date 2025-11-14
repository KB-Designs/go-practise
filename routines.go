package main

import (
	"fmt"
	"time"
)

func saymyname(name string) {
	fmt.Println("my name is", name)
	// time.Sleep(time.Second * 1)
}

func myposition(position int) {
	fmt.Println("my position is", position)
}

func main() {

	go saymyname("kimani")
	myposition(001)

	fmt.Println("yeah.. thats me")
	time.Sleep(time.Second * 2)

}

/*ch := make(chan int)

	years<-ch

	intenyears := years + 10

	fmt.Println(intenyears)

}

func getyears(yearborn int) int {
	totalyears := 2025 - yearborn
	ch <- totalyears
	return ch
}*/
