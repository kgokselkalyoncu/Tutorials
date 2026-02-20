package main

import "fmt"

func main() {

	// For loop with condition
	x := 0
	for x < 5 {
		fmt.Println("For loop iteration:", x)
		x++
	}

	fmt.Println("-------")

	// For loop with initialization, condition, and post statement
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	fmt.Println("-------")

	// Infinite loop (commented out to prevent execution)
	// for {
	// 	fmt.Println("This is an infinite loop")
	// }

	// For loop with array
	arr := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array element at index %d: %d\n", i, arr[i])
	}

	fmt.Println("-------")

	// For loop with range
	slice := []string{"Go", "is", "awesome"}

	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	fmt.Println("-------")

	// For loop with range and ignoring index
	for _, value := range slice {
		fmt.Printf("Value: %s\n", value)
	}

	fmt.Println("-------")

	// For loop with break statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop at i =", i)
			break
		}
		fmt.Println("For loop iteration:", i)
	}

	fmt.Println("-------")

	// For loop with continue statement
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Odd number:", i)
	}

}
