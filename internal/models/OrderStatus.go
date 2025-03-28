package models

type OrderStatus string

const (
	Pending   OrderStatus = "pending"
	Confirmed OrderStatus = "confirmed"
	Delivered OrderStatus = "delivered"
	Canceled  OrderStatus = "canceled"
)
