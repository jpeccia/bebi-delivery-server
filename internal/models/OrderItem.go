package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	SubTotal  float64 `json:"subtotal"`
}
