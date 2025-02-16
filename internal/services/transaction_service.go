package services

import (
	"merch-shop/internal/repository"
	"merch-shop/pkg/errors"
)

func SendCoins(senderID, receiverID string, amount int64) error {
	if amount <= 0 {
		return errors.ErrInvalidInput
	}

	return repository.DB.Transaction(func(tx repository.Tx) error {
		sender, err := repository.GetUserByID(senderID)
		if err != nil {
			return errors.ErrUserNotFound
		}

		receiver, err := repository.GetUserByID(receiverID)
		if err != nil {
			return errors.ErrUserNotFound
		}

		if sender.Balance < amount {
			return errors.ErrInsufficientFunds
		}

		sender.Balance -= amount
		receiver.Balance += amount

		if err := repository.UpdateUserBalance(sender); err != nil {
			return err
		}

		if err := repository.UpdateUserBalance(receiver); err != nil {
			return err
		}

		err = repository.SaveTransaction(senderID, receiverID, amount)
		if err != nil {
			return err
		}

		return nil
	})
}
