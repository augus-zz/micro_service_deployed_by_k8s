package model

import (
// "github.com/jinzhu/gorm"
)

type Customer struct {
	// gorm.Model
	ID      uint
	Name    string
	StoreID uint
	Store   Store
	UserID  uint
	User    User
}
