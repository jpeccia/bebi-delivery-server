package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	username string `json:"username"`
	password string `json:"password"`
	phone string `json:"phone"     gorm:"unique"`
	address string `json:"address"` 
	role Role `json:"role" gorm:"default:customer`
}