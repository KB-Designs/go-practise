package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    // Declare the atomic integer
    var counter atomic.Int32
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // Safely increment the counter by 1
            counter.Add(1)
        }()
    }

    wg.Wait()

    // Safely load the final value
    fmt.Printf("Final Counter Value: %d\n", counter.Load()) // Output: Final Counter Value: 1000
}
