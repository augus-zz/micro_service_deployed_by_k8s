package model

import (
	"testing"
)

func TestCreateCustomer(t *testing.T) {
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

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "hunan, changsha",
	}
	DB.Create(&store)

	customer := Customer{
		ID:      1,
		Name:    "augus",
		StoreID: store.ID, UserID: user.ID,
	}
	DB.Create(&customer)

	c := &Customer{ID: customer.ID}
	if err := DB.Preload("Store").Preload("User").Find(&c).Error; err != nil {
		t.Errorf("can not find customer, id: %v", c.ID)
	}

	if c.ID != customer.ID || c.Name != customer.Name {
		t.Errorf("can not find customer, id: %v", customer.ID)
	}

	if c.User.ID != user.ID || c.Store.ID != store.ID {
		t.Errorf("can not read association relation")
	}

	t.Logf("customer.User: %v", c.User)
	t.Logf("customer.Store: %v", c.Store)
}

func TestUpdateCustomer(t *testing.T) {
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

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "hunan, changsha",
	}
	DB.Create(&store)

	customer := Customer{
		ID:      1,
		Name:    "augus",
		StoreID: store.ID,
		UserID:  user.ID,
	}
	DB.Create(&customer)

	c := Customer{ID: customer.ID}
	if err := DB.Preload("Store").Preload("User").Find(&c).Error; err != nil {
		t.Errorf("can not find customer, id: %v", c.ID)
	}

	customer.Name = "Augus.zou"
	DB.Save(&customer)
	c = Customer{ID: customer.ID}
	if err := DB.Preload("Store").Preload("User").Find(&c).Error; err != nil {
		t.Errorf("can not find customer, id: %v", c.ID)
	}

	if c.ID != customer.ID || c.Name != customer.Name {
		t.Errorf("failed to update customer, id: %v", c.ID)
	}
}

func TestUpdateCustomerWithAssociation(t *testing.T) {
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

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "hunan, changsha",
	}
	DB.Create(&store)

	customer := Customer{
		ID:      1,
		Name:    "augus",
		StoreID: store.ID,
		UserID:  user.ID,
	}
	DB.Create(&customer)

	c := Customer{ID: customer.ID}
	if err := DB.Preload("Store").Preload("User").Find(&c).Error; err != nil {
		t.Errorf("can not find customer, id: %v", c.ID)
	}

	c.Name = "Augus.zou"
	c.User.Email = "augus.zou@gmail.com"
	c.Store.Title = "augus.zou's store"
	DB.Save(&c) //gorm:save_associations default true
	customer = Customer{ID: customer.ID}
	if err := DB.Preload("Store").Preload("User").Find(&customer).Error; err != nil {
		t.Errorf("can not find customer, id: %v", customer.ID)
	}

	if c.ID != customer.ID || c.Name != customer.Name {
		t.Errorf("failed to update customer, id: %v", customer.ID)
	}

	var u User
	DB.First(&u, user.ID)
	if u.ID != customer.User.ID || u.Email != customer.User.Email || c.User.Email != customer.User.Email {
		t.Errorf("failed to update customer.user, id: %v", customer.User.ID)
	}

	var s Store
	DB.First(&s, store.ID)
	if s.ID != customer.Store.ID || s.Title != customer.Store.Title || c.Store.Title != customer.Store.Title {
		t.Errorf("failed to update customer.store, id: %v", customer.Store.ID)
	}
}

func TestDeleteCustomer(t *testing.T) {
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

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "hunan, changsha",
	}
	DB.Create(&store)

	customer := Customer{
		ID:      1,
		Name:    "augus",
		StoreID: store.ID,
		UserID:  user.ID,
	}
	DB.Create(&customer)

	c := Customer{ID: customer.ID}
	if err := DB.Preload("Store").Preload("User").Find(&c).Error; err != nil {
		t.Errorf("can not find customer, id: %v", c.ID)
	}

	DB.Delete(&customer)

	if !DB.First(&c, customer.ID).RecordNotFound() {
		t.Errorf("failed to delete customer, id: %v", c.ID)
	}

	var u User
	if DB.First(&u, user.ID).RecordNotFound() {
		t.Errorf("delete customer with user, customer.id: %v, user.id: %v", c.ID, user.ID)
	}

	var s Store
	if DB.First(&s, store.ID).RecordNotFound() {
		t.Errorf("delete customer with store, customer.id: %v, store.id: %v", c.ID, store.ID)
	}
}

func TestDeleteCustomerWithAssociation(t *testing.T) {
	clearDB()
	defer clearDB()

	t.Skip("will not remove relation use Delete method")

	user := User{
		ID:    1,
		Name:  "zouqilin",
		Phone: "12345678",
		Email: "mr.zouqilin@gmail.com",
		Role:  ROLE_GOD,
	}

	DB.Create(&user)

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "hunan, changsha",
	}
	DB.Create(&store)

	customer := Customer{
		ID:      1,
		Name:    "augus",
		StoreID: store.ID,
		UserID:  user.ID,
	}
	DB.Create(&customer)

	c := Customer{ID: customer.ID}
	if err := DB.Preload("Store").Preload("User").Find(&c).Error; err != nil {
		t.Errorf("can not find customer, id: %v", c.ID)
	}

	DB.Delete(&c)

	if !DB.First(&c, customer.ID).RecordNotFound() {
		t.Errorf("failed to delete customer, id: %v", c.ID)
	}

	var u User
	if DB.First(&u, user.ID).RecordNotFound() {
		t.Errorf("delete customer with user, customer.id: %v, user.id: %v", c.ID, user.ID)
	}

	var s Store
	if DB.First(&s, store.ID).RecordNotFound() {
		t.Errorf("delete customer with store, customer.id: %v, store.id: %v", c.ID, store.ID)
	}
}
