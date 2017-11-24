package model

import (
// "github.com/jinzhu/gorm"
)

type Customer struct {
	// gorm.Model
	ID      uint32
	Name    string
	StoreID uint32
	Store   Store
	UserID  uint32
	User    User
}
