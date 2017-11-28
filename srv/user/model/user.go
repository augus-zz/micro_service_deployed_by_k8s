package model

import (
// "github.com/jinzhu/gorm"
)

const (
	ROLE_CUSTOMER = iota
	ROLE_STORE_STUFF
	ROLE_STORE_MANAGER
	ROLE_STORE_ADMIN
	ROLE_GOD
)

var (
	USER_ROLE_CUSTOMER      string = "customer"
	USER_ROLE_STORE_STUFF   string = "stuff"
	USER_ROLE_STORE_MANAGER string = "manager"
	USER_ROLE_STORE_ADMIN   string = "admin"
	USER_ROLE_GOD           string = "god"
)

var USER_ROLES []string = []string{
	USER_ROLE_CUSTOMER,
	USER_ROLE_STORE_STUFF,
	USER_ROLE_STORE_MANAGER,
	USER_ROLE_STORE_ADMIN,
	USER_ROLE_GOD,
}

type User struct {
	// gorm.Model
	ID    uint32
	Name  string
	Phone string
	Email string
	Role  uint32
}

func (user *User) UserRole() string {
	if user.Role > ROLE_GOD {
		return ""
	}
	return USER_ROLES[user.Role]
}
