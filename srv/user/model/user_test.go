package model

import (
	"micro_service_deployed_by_k8s/shared/db"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ClearDB()
	defer ClearDB()

	user := User{
		ID:    1,
		Name:  "zouqilin",
		Phone: "12345678",
		Email: "mr.zouqilin@gmail.com",
		Role:  ROLE_GOD,
	}

	db.DB.Create(&user)

	var u User
	db.DB.First(&u, user.ID)

	if u.ID != user.ID || u.Name != user.Name {
		t.Errorf("failed to create user")
	}
}

func TestUpdateUser(t *testing.T) {
	ClearDB()
	defer ClearDB()

	user := User{
		ID:    1,
		Name:  "zouqilin",
		Phone: "12345678",
		Email: "mr.zouqilin@gmail.com",
		Role:  ROLE_GOD,
	}

	db.DB.Create(&user)

	user.Phone = "87654321"

	db.DB.Save(&user)

	var u User
	db.DB.First(&u, user.ID)

	if u.ID != user.ID || u.Phone != user.Phone {
		t.Errorf("failed to update user")
	}
}

func TestDeleteUser(t *testing.T) {
	ClearDB()
	defer ClearDB()

	user := User{
		ID:    1,
		Name:  "zouqilin",
		Phone: "12345678",
		Email: "mr.zouqilin@gmail.com",
		Role:  ROLE_GOD,
	}

	db.DB.Create(&user)

	var u User
	if err := db.DB.First(&u, user.ID).Error; err != nil {
		t.Errorf("failed to create user")
	}

	id := user.ID
	db.DB.Delete(&user)

	if !db.DB.First(&u, id).RecordNotFound() {
		t.Errorf("failed to delete create user")
	}
}
