package repository

import (
	"merch-shop/internal/db"
	"merch-shop/internal/models"
)

func GetItemByName(name string) (*models.Merch, error) {
	var item models.Merch
	err := db.DB.Get(&item, "SELECT * FROM merch WHERE name = $1 AND deleted_at IS NULL", name)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func GetAllItems() ([]models.Merch, error) {
	var items []models.Merch
	err := db.DB.Select(&items, "SELECT * FROM merch WHERE deleted_at IS NULL ORDER BY name")
	return items, err
}
