package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	NumPedido  int         `json:"num_pedido"`
	CustomerID uint        `json:"customer_id"`
	StoreId    uint        `json:"store_id"`
	Status     OrderStatus `json:"status" gorm:"default:'pending'"`
	TotalPrice float64     `json:"total_price"`
}
