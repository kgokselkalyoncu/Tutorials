package main

import (
	"_packages_and_modules/student"
	"_packages_and_modules/teacher"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	sayHello("Kartal")

	result := add(5, 3)
	fmt.Printf("The result of adding 5 and 3 is: %d\n", result)

	student.SayHello("Alice")
	grade := student.CalcGrade(85)
	fmt.Printf("The grade for score 85 is: %s\n", grade)

	teacher.SayHello("Mr. Smith")
}
