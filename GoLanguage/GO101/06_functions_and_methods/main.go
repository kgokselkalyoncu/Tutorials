package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello, ", name)
}

func main() {

	sayHello("Kartal")
	sayHello("Göksel")

	fmt.Println("Functions and Methods in Go")

	// Function example
	result := add(5, 3)
	fmt.Println("Result of add function:", result)

	areaRect := calculateAreaRect(4.5, 2.0)
	fmt.Printf("Area of rectangle: %.2f\n", areaRect)

	areaCircle := calculateAreaCircle(5.0)
	fmt.Printf("Area of circle (using radius as f1 and pi as f2): %.2f\n", areaCircle)

	areaSquare := calculateAreaSquare(4.0)
	fmt.Printf("Area of square: %.2f\n", areaSquare)

}

func calculateAreaRect(f1, f2 float64) float64 {
	return f1 * f2
}

func calculateAreaCircle(r float64) float64 {
	return 3.14159 * r * r
}

func calculateAreaSquare(a float64) float64 {
	return a * a
}

func add(i1, i2 int) int {
	return i1 + i2 // Return type is 'any' to allow for flexibility, but it can be changed to 'int' if only integers are expected + panic("unimplemented")
}
