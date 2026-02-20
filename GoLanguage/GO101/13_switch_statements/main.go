package main

import (
	product "_switch_statements/models"
	"fmt"
)

var _product *product.Product

func Create() {
	// Initialize a product for demonstration purposes.

	var name string
	var description string
	var category string
	var price float64
	var stock int

	fmt.Println("Creating a new product...")
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Print("Description: ")
	fmt.Scanln(&description)
	fmt.Print("Category: ")
	fmt.Scanln(&category)
	fmt.Print("Price: ")
	fmt.Scanln(&price)
	fmt.Print("Stock: ")
	fmt.Scanln(&stock)

	p := product.NewProduct(name, description, category, price, stock)
	_product = p
	fmt.Println("Product created successfully:", p)
}

func List() {
	// List all products.
	fmt.Println("Listing all products:")
	products := product.GetAllProducts()
	for _, p := range products {
		fmt.Println(p)
	}
}

func Get() {
	// Get a product by ID.
	var id int
	fmt.Println("Enter product ID to get:")
	fmt.Scanln(&id)
	p := product.GetProductById(id)
	if p != nil {
		fmt.Println("Product found:", p)
	} else {
		fmt.Println("Product not found.")
	}
}

func Update() {
	// Update the stock of a product.
	var id int
	fmt.Println("Enter product ID to update stock:")
	fmt.Scanln(&id)
	p := product.GetProductById(id)
	if p != nil {
		var amount int
		fmt.Println("Enter stock amount to add (use negative value to reduce):")
		fmt.Scanln(&amount)
		p.UpdateStock(amount)
		fmt.Println("Stock updated successfully. New product details:", p)
	} else {
		fmt.Println("Product not found.")
	}
}

func Delete() {
	// Delete a product by ID.
	var id int
	fmt.Println("Enter product ID to delete:")
	fmt.Scanln(&id)
	p := product.GetProductById(id)
	if p != nil {
		if product.DeleteProductById(id) {
			fmt.Println("Product deleted successfully.")
		} else {
			fmt.Println("Failed to delete product.")
		}
	} else {
		fmt.Println("Product not found.")
	}
}

func mainMenu() {
	for {
		fmt.Println("\nProduct Management System")
		fmt.Println("1. Create Product")
		fmt.Println("2. List Products")
		fmt.Println("3. Get Product by ID")
		fmt.Println("4. Update Product Stock")
		fmt.Println("5. Delete Product")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			Create()
		case 2:
			List()
		case 3:
			Get()
		case 4:
			Update()
		case 5:
			Delete()
		case 6:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func switchStatement() {
	// switch statement is a control flow statement that allows you to execute a block of code based on the value of an expression.
	// It is similar to if-else statements but is more concise and easier to read when you have multiple conditions to check.

	// Basic switch statement
	var text string
	fmt.Println("Enter a day of the week:")
	fmt.Scanln(&text)

	switch text {
	case "Monday":
		fmt.Println("It's the start of the week!")
	case "Tuesday":
		fmt.Println("It's the second day of the week!")
	case "Wednesday":
		fmt.Println("It's the middle of the week!")
	case "Thursday":
		fmt.Println("It's almost the end of the week!")
	case "Friday":
		fmt.Println("It's the end of the week!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's not a regular day.")
	}
}

func main() {
	// switchStatement()
	mainMenu()
}
