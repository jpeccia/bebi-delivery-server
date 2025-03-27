package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name string `json:"name"`
	OwnerID uint `json"owner_id"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Rating int `json:"rating"`
	Address string `json:"address"`
}