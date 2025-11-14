package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string) // Unbuffered channel

	go func() {
		fmt.Println("Goroutine sending message...")
		messages <- "Hello from a channel!" // Send a message
		fmt.Println("Goroutine sent message.")
	}()

	fmt.Println("Main goroutine waiting for message...")
	msg := <-messages // Receive a message
	fmt.Println("Main goroutine received:", msg)
	time.Sleep(10 * time.Millisecond) // Give time for goroutine to print its last message
}
