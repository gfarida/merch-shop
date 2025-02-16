package repository

import (
	"merch-shop/internal/db"
	"merch-shop/internal/models"
)

func AddToInventory(userID, merchID string, quantity int) error {
	var inventoryItem models.InventoryItem
	err := db.DB.Get(&inventoryItem, `
		SELECT * FROM inventory 
		WHERE user_id = $1 AND merch_id = $2`, userID, merchID)
	if err == nil {
		_, err := db.DB.Exec(`
			UPDATE inventory 
			SET quantity = quantity + $1 
			WHERE user_id = $2 AND merch_id = $3`, quantity, userID, merchID)
		return err
	}

	_, err = db.DB.Exec(`
		INSERT INTO inventory (user_id, merch_id, quantity) 
		VALUES ($1, $2, $3)`, userID, merchID, quantity)
	return err
}
