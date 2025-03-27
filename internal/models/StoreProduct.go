package models

import (
	"gorm.io/gorm"
)

type StoreProduct struct {
	gorm.Model
	StoreId uint `json:"store_id"`
	ProductId uint `json:"product_id"`
	Quantity int `json:"quantity"`
	
	Store     Store     `gorm:"foreignKey:StoreID" json:"store"`
    Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
}