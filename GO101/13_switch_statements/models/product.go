package models

import "fmt"

// Product represents a product with a name and price.
type Product struct {
	Id          int
	Name        string
	Description string
	Category    string
	Price       float64
	Stock       int
}

var _idCounter int = 1
var products []Product

// NewProduct creates a new product with the given details.
func NewProduct(name, description, category string, price float64, stock int) *Product {
	product := Product{
		Id:          _idCounter,
		Name:        name,
		Description: description,
		Category:    category,
		Price:       price,
		Stock:       stock,
	}

	_idCounter++
	products = append(products, product)
	return &product
}

// UpdateStock updates the stock of the product by the given amount.
func (p *Product) UpdateStock(amount int) { p.Stock += amount }

// IsInStock checks if the product is in stock.
func (p *Product) IsInStock() bool { return p.Stock > 0 }

// GetPrice returns the price of the product.
func (p *Product) GetPrice() float64 { return p.Price }

// GetCategory returns the category of the product.
func (p *Product) GetCategory() string { return p.Category }

// GetDescription returns the description of the product.
func (p *Product) GetDescription() string { return p.Description }

// GetName returns the name of the product.
func (p *Product) GetName() string { return p.Name }

// String returns a string representation of the product.
func (p *Product) String() string {
	return p.Name + " - " + p.Description + " (" + p.Category + ") - $" + fmt.Sprintf("%.2f", p.Price) + " - Stock: " + fmt.Sprintf("%d", p.Stock)
}

// GetAllProducts returns a slice of all products.
func GetAllProducts() []Product { return products }

// GetProductById returns a product by its ID.
func GetProductById(id int) *Product {
	for _, product := range products {
		if product.Id == id {
			return &product
		}
	}
	return nil
}

// GetProductsByCategory returns a slice of products in the given category.
func GetProductsByCategory(category string) []Product {
	var categoryProducts []Product
	for _, product := range products {
		if product.Category == category {
			categoryProducts = append(categoryProducts, product)
		}
	}
	return categoryProducts
}

// GetProductsInPriceRange returns a slice of products within the given price range.
func GetProductsInPriceRange(minPrice, maxPrice float64) []Product {
	var priceRangeProducts []Product
	for _, product := range products {
		if product.Price >= minPrice && product.Price <= maxPrice {
			priceRangeProducts = append(priceRangeProducts, product)
		}
	}
	return priceRangeProducts
}

// DeleteProductById deletes a product by its ID.
func DeleteProductById(id int) bool {
	for i, product := range products {
		if product.Id == id {
			products = append(products[:i], products[i+1:]...)
			return true
		}
	}
	return false
}
