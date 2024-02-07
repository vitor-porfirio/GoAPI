package entity

import "github.com/google/uuid"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	// Category *Category
	CategoryID string `json:"category_id"`
	ImageURL   string `json:"image_url"`
}

// NewProduct creates a new instance of the Product struct with the provided parameters.
// It generates a new UUID for the ID field and initializes the other fields with the given values.
func NewProduct(name, description string, price float64, categoryID, imageURL string) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		ImageURL:    imageURL,
	}
}
