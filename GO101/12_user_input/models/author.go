package models

import (
	"fmt"
	"strings"
)

type Author struct {
	Name  string
	Age   int
	Books []string
}

// CreateAuthor prompts the user for author details and returns an Author struct
func CreateAuthor() Author {

	var name string
	var age int
	var books []string
	fmt.Print("Enter author's name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter author's age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter author's books (comma separated): ")
	var booksInput string
	fmt.Scanln(&booksInput)
	books = strings.Split(booksInput, ",")

	return Author{
		Name:  name,
		Age:   age,
		Books: books,
	}
}

// String returns a string representation of the Author struct
func (a Author) String() string {
	return fmt.Sprintf("Author: %s, Age: %d, Books: %v", a.Name, a.Age, a.Books)
}
