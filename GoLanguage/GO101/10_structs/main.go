package main

import (
	student "_structs/models"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Lessons []string
}

func main() {

	fmt.Println("Structs in Go")

	p1 := Person{Name: "Alice", Age: 30}
	p1.Lessons = []string{"Math", "Science", "History"}

	fmt.Printf("Person 1: %+v\n", p1)

	p2 := Person{Name: "Bob", Age: 25}
	p2.Lessons = []string{"English", "Art"}

	fmt.Println("Person 2:", p2)
	fmt.Printf("Person 2 Name: %s, Age: %d, Lessons: %v\n", p2.Name, p2.Age, p2.Lessons)

	p3 := Person{Name: "Charlie"}
	p3.Age = 40
	p3.Lessons = []string{"Music", "Physical Education"}

	fmt.Printf("Person 3: %+v\n", p3)
	fmt.Printf("Person 3 Name: %s, Age: %d, Lessons: %v\n", p3.Name, p3.Age, p3.Lessons)

	student1 := student.NewStudent("David", 20, []string{"Programming", "Data Structures"})
	student1 = student.AddLesson(student1, "Databases")
	fmt.Printf("Student 1: %+v\n", student1)

	student2 := student.CreateStudentEmpty()
	student2.Name = "Eve"
	student2.Age = 22
	student2 = student.AddLesson(student2, "Algorithms")
	student2 = student.AddLesson(student2, "Operating Systems")
	fmt.Printf("Student 2: %+v\n", student2)
}
