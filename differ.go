package main

import "fmt"

func openFile(name string) {
    fmt.Println("Opening", name)
}

func closeFile(name string) {
    fmt.Println("Closing", name)
}

func main() {
    openFile("file1")
    defer closeFile("file1")

    openFile("file2")
    defer closeFile("file2")

    fmt.Println("Reading files...")
}
