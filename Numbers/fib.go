// Fibonacci
package main

import "fmt"

func main() {
    fmt.Println(fib(20))
}

func fib (x int) int {
    if x <= 1 {
        return x
    } else {
        return fib(x - 1) + fib(x - 2)
    }
}