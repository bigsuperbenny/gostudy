package main

import "fmt"

func main() {
    // Print a simple message
    fmt.Println("Hello, World!")

    // Declare and initialize some variables
    name := "Gopher"
    age := 10

    // Using string formatting
    fmt.Printf("My name is %s and I am %d years old\n", name, age)

    // Simple addition
    result := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
}

// Simple function that adds two numbers
func add(a, b int) int {
    return a + b
}