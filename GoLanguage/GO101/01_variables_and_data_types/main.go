package main

import "fmt"

func main() {

	// Variable declaration and initialization
	var name string = "Go101"

	// Short variable declaration (type inference)
	var surname = "Gopher"
	surname = "Golang"

	var school string
	school = "Go101 School"

	fmt.Println("Hello, " + name + " " + surname + "! You are enrolled in " + school + ".")

	// Multiple variable declaration
	var age, year int
	age = 5
	year = 2024

	fmt.Printf("You are %d years old in the year %d.\n", age, year)

	// Integer variable declaration
	var height int = 170
	fmt.Printf("Your height is %d cm.\n", height)

	// Float variable declaration
	var pi float64 = 3.14159
	fmt.Printf("The value of pi is %.2f.\n", pi)

	// Integer and Float variable declaration
	var radius int = 10
	var area float64 = pi * float64(radius*radius)
	fmt.Printf("The area of a circle with radius %d is %.2f.\n", radius, area)

	// Boolean variable declaration
	var isGoFun bool = true
	fmt.Printf("Is Go fun? %t\n", isGoFun)
}
