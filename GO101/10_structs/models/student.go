package models

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

func AddLesson(s Student, lesson string) Student {
	s.Lessons = append(s.Lessons, lesson)
	return s
}
