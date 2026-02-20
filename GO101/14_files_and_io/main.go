package main

import (
	"fmt"
	"os"
)

func main() {
	// Example: Reading and writing files
	filename := "example.txt"
	
	// Write to file
	data := []byte("Hello, Go!\n")
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	
	// Read from file
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	fmt.Println("File content:", string(content))
}