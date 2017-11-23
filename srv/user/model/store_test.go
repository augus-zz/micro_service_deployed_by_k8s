package model

import (
	"testing"
)

func TestCreateStore(t *testing.T) {
	clearDB()
	defer clearDB()

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "hunan, changsha",
	}

	DB.Create(&store)

	var s Store
	DB.First(&s, store.ID)

	if s.ID != store.ID || s.Title != store.Title {
		t.Errorf("failed to create store")
	}
}

func TestUpdateStore(t *testing.T) {
	clearDB()
	defer clearDB()

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "changsha, hunan",
	}

	DB.Create(&store)

	store.Address = "shenzhen, guangdong"

	DB.Save(&store)

	var s Store
	DB.First(&s, store.ID)

	if s.ID != store.ID || s.Address != store.Address {
		t.Errorf("failed to update store")
	}
}

func TestDeleteStore(t *testing.T) {
	clearDB()
	defer clearDB()

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "changsha, hunan",
	}

	DB.Create(&store)

	var s Store
	if err := DB.First(&s, store.ID).Error; err != nil {
		t.Errorf("failed to create store")
	}

	id := store.ID
	DB.Delete(&store)

	if !DB.First(&s, id).RecordNotFound() {
		t.Errorf("failed to delete create store")
	}
}
