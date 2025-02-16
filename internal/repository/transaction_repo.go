package repository

import (
	"merch-app/internal/db"
	"merch-app/internal/models"
)

func SaveTransaction(senderID, receiverID string, amount int64, transactionTypeID int) error {
	_, err := db.DB.Exec(`
		INSERT INTO transactions (id, sender_id, receiver_id, transaction_type_id, amount, created_at)
		VALUES (gen_random_uuid(), $1, $2, $3, $4, now())`,
		senderID, receiverID, transactionTypeID, amount)
	return err
}

func GetUserTransactions(userID string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := db.DB.Select(&transactions, `
		SELECT id, sender_id, receiver_id, transaction_type_id, amount, created_at
		FROM transactions 
		WHERE sender_id = $1 OR receiver_id = $1
		ORDER BY created_at DESC`, userID)
	return transactions, err
}
