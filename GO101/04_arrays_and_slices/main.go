package main

import "fmt"

func main() {

	// Arrays
	var arr [5]int
	fmt.Println("Array:", arr)
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("Updated Array:", arr)

	fmt.Println("arr[2]:", arr[2])

	// Slices
	slice := []int{1, 2, 3}
	fmt.Println("Slice:", slice)
	slice = append(slice, 4)
	fmt.Println("Updated Slice:", slice)

	// Accessing elements
	fmt.Println("slice[1]:", slice[1])

	// Slicing
	subSlice := slice[1:3]
	fmt.Println("Sub-slice:", subSlice)

	// Length and Capacity
	fmt.Println("Slice Length:", len(slice))
	fmt.Println("Slice Capacity:", cap(slice))

}
