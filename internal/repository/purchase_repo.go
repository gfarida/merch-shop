package repository

import (
	"merch-app/internal/db"
	"merch-app/internal/models"
)

func SavePurchase(userID, merchID string) error {
	_, err := db.DB.Exec(`
		INSERT INTO purchases (id, user_id, merch_id, created_at)
		VALUES (gen_random_uuid(), $1, $2, now())`, userID, merchID)
	return err
}

func GetUserPurchases(userID string) ([]models.Purchase, error) {
	var purchases []models.Purchase
	err := db.DB.Select(&purchases, `
		SELECT id, user_id, merch_id, created_at
		FROM purchases
		WHERE user_id = $1
		ORDER BY created_at DESC`, userID)
	return purchases, err
}
