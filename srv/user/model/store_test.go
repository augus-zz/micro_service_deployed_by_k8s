package model

import (
	"micro_service_deployed_by_k8s/shared/db"
	"testing"
)

func TestCreateStore(t *testing.T) {
	ClearDB()
	defer ClearDB()

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "hunan, changsha",
	}

	db.DB.Create(&store)

	var s Store
	db.DB.First(&s, store.ID)

	if s.ID != store.ID || s.Title != store.Title {
		t.Errorf("failed to create store")
	}
}

func TestUpdateStore(t *testing.T) {
	ClearDB()
	defer ClearDB()

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "changsha, hunan",
	}

	db.DB.Create(&store)

	store.Address = "shenzhen, guangdong"

	db.DB.Save(&store)

	var s Store
	db.DB.First(&s, store.ID)

	if s.ID != store.ID || s.Address != store.Address {
		t.Errorf("failed to update store")
	}
}

func TestDeleteStore(t *testing.T) {
	ClearDB()
	defer ClearDB()

	store := Store{
		ID:      1,
		Title:   "zouqilin's store",
		Address: "changsha, hunan",
	}

	db.DB.Create(&store)

	var s Store
	if err := db.DB.First(&s, store.ID).Error; err != nil {
		t.Errorf("failed to create store")
	}

	id := store.ID
	db.DB.Delete(&store)

	if !db.DB.First(&s, id).RecordNotFound() {
		t.Errorf("failed to delete create store")
	}
}
