package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
    Username string `json:"username"`
    Password string `json:"password"`
    Phone    string `json:"phone" gorm:"unique"`
    Address  string `json:"address"`
	Role Role `json:"role" gorm:"default:customer`
}