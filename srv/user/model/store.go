package model

import (
// "github.com/jinzhu/gorm"
)

type Store struct {
	// gorm.Model
	ID      uint
	Title   string
	Address string
}
