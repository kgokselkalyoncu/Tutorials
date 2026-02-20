package main

import "fmt"

func main() {

	// Print a string to the console without a newline at the end
	fmt.Print("Hello, World!")
	fmt.Print("Hello, World!\n") // Manually add a newline character

	fmt.Print("This new line")

	// Print a string to the console with a newline at the end
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")

	var name string = "Alice"
	var age int = 30
	var salary float64 = 75000.87

	// Use fmt.Println to print multiple values with a space in between
	fmt.Println("Hello ", name, age)

	// Use fmt.Printf to format the output with placeholders
	fmt.Printf("Name: %v, Age: %v\n", name, age)

	// %v is a placeholder for any value, and \n adds a newline at the end of the output
	// %s is a placeholder for a string
	// %d is a placeholder for an integer
	// %q is a placeholder for a quoted string
	// %T is a placeholder for the type of the value
	// %f is a placeholder for a floating-point number

	fmt.Printf("Name: %s, Age: %d \n", name, age)
	fmt.Printf("Salary: %f \n", salary)   // Format salary to any decimal places
	fmt.Printf("Salary: %.2f \n", salary) // Format salary to 2 decimal places
	fmt.Printf("Salary: %.1f \n", salary) // Format salary to 1 decimal places
	fmt.Printf("Salary: %.0f \n", salary) // Format salary to 0 decimal places
	fmt.Printf("Quoted Name: %q \n", name)
	fmt.Printf("Type of name: %T \n", name)

	// Use fmt.Sprintf to format a string and store it in a variable
	// formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
	var formattedString string = fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(formattedString)

}
