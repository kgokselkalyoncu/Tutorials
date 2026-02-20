package main

import "fmt"

// func changeValue(x int) {
// 	x = 10
// }

func changeValue(x *int) {

	fmt.Println("Address of x:", x)
	fmt.Println("Value of x:", *x)

	*x = 10
}

func main() {
	y := 5
	fmt.Println("Value of y before changeValue:", y)
	fmt.Println("Address of y:", &y)
	// changeValue(y) // This will not change the value of y in main
	changeValue(&y) // This will change the value of y in main
	fmt.Println("Value of y after changeValue:", y)
	fmt.Println("Address of y:", &y)
	fmt.Println("Hello, World!")
}
