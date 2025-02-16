package repository

import (
	"merch-app/internal/db"
	"merch-app/internal/models"
)

func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL", userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE username = $1 AND deleted_at IS NULL", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserBalance(userID string, newBalance int64) error {
	_, err := db.DB.Exec("UPDATE users SET balance = $1, updated_at = now() WHERE id = $2", newBalance, userID)
	return err
}
