package main

import "fmt"

func main() {

	// Create a map to store menu items and their prices
	menu := make(map[string]float64)
	menu["coffee"] = 2.50
	menu["tea"] = 1.75
	menu["sandwich"] = 5.00

	fmt.Println(menu)

	fmt.Println(menu["coffee"]) // Accessing a value by key

	menu["cake"] = 3.00 // Adding a new key-value pair

	fmt.Println(menu)

	menu["tea"] = 1.50 // Updating an existing value

	fmt.Println(menu)

	delete(menu, "sandwich") // Deleting a key-value pair

	fmt.Println(menu)

	fmt.Println("---------")

	// Alternatively, we can initialize the map with values directly
	menu2 := map[string]float64{
		"coffee":   2.50,
		"tea":      1.75,
		"sandwich": 5.00,
	}

	fmt.Println(menu2)

	// Another example: a phone book mapping names to phone numbers
	phoneBook := map[string]string{
		"Alice":   "123-456-7890",
		"Bob":     "987-654-3210",
		"Charlie": "555-555-5555",
	}

	fmt.Println(phoneBook)

	// Iterating over the map
	for item, price := range menu2 {
		fmt.Printf("%s: $%.2f\n", item, price)
	}

	for name, number := range phoneBook {
		fmt.Printf("%s: %s\n", name, number)
	}
}
