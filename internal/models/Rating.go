package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	StoreID     uint   `json:"store_id"`
	UserID      uint   `json:"user_id"`
}
