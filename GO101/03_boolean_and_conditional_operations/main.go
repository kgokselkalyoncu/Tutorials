package main

import "fmt"

func main() {

	// Boolean values
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println("isTrue:", isTrue)
	fmt.Println("isFalse:", isFalse)

	// Conditional statements
	if isTrue {
		fmt.Println("This is true!")
	} else {
		fmt.Println("This is false!")
	}

	// Logical operators
	var a bool = true
	var b bool = false

	fmt.Println("a AND b:", a && b) // Logical AND
	fmt.Println("a OR b:", a || b)  // Logical OR
	fmt.Println("NOT a:", !a)       // Logical NOT

	// Comparison operators
	var x int = 10
	var y int = 20

	fmt.Println("x == y:", x == y) // Equal to
	fmt.Println("x != y:", x != y) // Not equal to
	fmt.Println("x > y:", x > y)   // Greater than
	fmt.Println("x < y:", x < y)   // Less than
	fmt.Println("x >= y:", x >= y) // Greater than or equal to
	fmt.Println("x <= y:", x <= y) // Less than or equal to
}
