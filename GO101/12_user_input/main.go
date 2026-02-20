package main

import (
	author "_user_input/models"
	"fmt"
)

func main() {
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)
	// fmt.Printf("Hello, %s!\n", name)

	author := author.CreateAuthor()
	fmt.Println(author.String())
}
