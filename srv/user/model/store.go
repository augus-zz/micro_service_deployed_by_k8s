package model

import (
// "github.com/jinzhu/gorm"
)

type Store struct {
	// gorm.Model
	ID      uint32
	Title   string
	Address string
}
