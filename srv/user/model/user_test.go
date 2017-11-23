package model

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	clearDB()
	defer clearDB()

	user := User{
		ID:    1,
		Name:  "zouqilin",
		Phone: "12345678",
		Email: "mr.zouqilin@gmail.com",
		Role:  ROLE_GOD,
	}

	DB.Create(&user)

	var u User
	DB.First(&u, user.ID)

	if u.ID != user.ID || u.Name != user.Name {
		t.Errorf("failed to create user")
	}
}

func TestUpdateUser(t *testing.T) {
	clearDB()
	defer clearDB()

	user := User{
		ID:    1,
		Name:  "zouqilin",
		Phone: "12345678",
		Email: "mr.zouqilin@gmail.com",
		Role:  ROLE_GOD,
	}

	DB.Create(&user)

	user.Phone = "87654321"

	DB.Save(&user)

	var u User
	DB.First(&u, user.ID)

	if u.ID != user.ID || u.Phone != user.Phone {
		t.Errorf("failed to update user")
	}
}

func TestDeleteUser(t *testing.T) {
	clearDB()
	defer clearDB()

	user := User{
		ID:    1,
		Name:  "zouqilin",
		Phone: "12345678",
		Email: "mr.zouqilin@gmail.com",
		Role:  ROLE_GOD,
	}

	DB.Create(&user)

	var u User
	if err := DB.First(&u, user.ID).Error; err != nil {
		t.Errorf("failed to create user")
	}

	id := user.ID
	DB.Delete(&user)

	if !DB.First(&u, id).RecordNotFound() {
		t.Errorf("failed to delete create user")
	}
}
