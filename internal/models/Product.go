package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
	Value float64 `json:"value"`
	ImageUrl string `json:"imageUrl"`
}
