package main

import "fmt"

// Define a struct type
type Person struct {
    Name string
    Age  int
}

// Define a method with a value receiver
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Define a method with a pointer receiver
func (p *Person) SetAge(newAge int) {
    p.Age = newAge
}

func main() {
    // Create an instance of the struct
    person1 := Person{Name: "Alice", Age: 30}

    // Call the method with a value receiver
    person1.Greet() // Output: Hello, my name is Alice and I am 30 years old.

    // Call the method with a pointer receiver
    person1.SetAge(31)
    person1.Greet() // Output: Hello, my name is Alice and I am 31 years old.
}
