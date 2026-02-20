package main

import (
	student "_receivers/models"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	student1 := student.NewStudent("David", 20, []string{"Programming", "Data Structures"})

	// Using the method with pointer receiver to modify the original Student instance
	student1.AddLesson("Algorithms")

	// Using the function to add a lesson (returns a new Student instance)
	student1 = student.AddLesson(student1, "Operating Systems")

	fmt.Printf("Student 1: %+v\n", student1)

	// Using the method with value receiver to display student information
	student1.DisplayStudent()

	// Using the function to display student information
	student.DisplayStudent(student1)

}
