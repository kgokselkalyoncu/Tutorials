package models

import "fmt"

type Student struct {
	Name    string
	Age     int
	Lessons []string
}

func CreateStudentEmpty() Student {
	return Student{
		Name:    "",
		Age:     0,
		Lessons: []string{},
	}
}

func NewStudent(name string, age int, lessons []string) Student {
	return Student{
		Name:    name,
		Age:     age,
		Lessons: lessons,
	}
}

// Function to add a lesson to a Student instance (returns a new Student with the added lesson)
func AddLesson(s Student, lesson string) Student {
	s.Lessons = append(s.Lessons, lesson)
	return s
}

// Method with pointer receiver to modify the original Student instance
func (s *Student) AddLesson(lesson string) {
	s.Lessons = append(s.Lessons, lesson)
}

// Method with value receiver to display student information (does not modify the Student instance)
func (s Student) DisplayStudent() {
	fmt.Printf("Name: %s, Age: %d, Lessons: %v\n", s.Name, s.Age, s.Lessons)
}

// Function to display student information (takes a Student instance as an argument)
func DisplayStudent(s Student) {
	fmt.Printf("Name: %s, Age: %d, Lessons: %v\n", s.Name, s.Age, s.Lessons)
}
