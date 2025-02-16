package services

import (
	"log/slog"
	"merch-shop/internal/repository"
	"merch-shop/pkg/errors"
	"merch-shop/pkg/logger"
)

func SendCoins(senderID, receiverID string, amount int64) error {
	if amount <= 0 {
		logger.Log().Error("Invalid amount for transaction", slog.Int64("amount", amount))
		return errors.ErrInvalidInput
	}

	return repository.DB.Transaction(func(tx repository.Tx) error {
		sender, err := repository.GetUserByID(senderID)
		if err != nil {
			logger.Log().Error("Sender not found", slog.String("sender_id", senderID))
			return errors.ErrUserNotFound
		}

		receiver, err := repository.GetUserByID(receiverID)
		if err != nil {
			logger.Log().Error("Receiver not found", slog.String("receiver_id", receiverID))
			return errors.ErrUserNotFound
		}

		if sender.Balance < amount {
			logger.Log().Error("Insufficient funds", slog.String("sender_id", senderID), slog.Int64("balance", sender.Balance), slog.Int64("amount", amount))
			return errors.ErrInsufficientFunds
		}

		sender.Balance -= amount
		receiver.Balance += amount

		// Обновляем баланс пользователей
		if err := repository.UpdateUserBalance(sender.ID, sender.Balance); err != nil {
			logger.Log().Error("Failed to update sender balance", slog.String("sender_id", senderID), slog.Int64("new_balance", sender.Balance))
			return err
		}

		if err := repository.UpdateUserBalance(receiver.ID, receiver.Balance); err != nil {
			logger.Log().Error("Failed to update receiver balance", slog.String("receiver_id", receiverID), slog.Int64("new_balance", receiver.Balance))
			return err
		}

		// Сохраняем транзакцию
		err = repository.SaveTransaction(senderID, receiverID, amount)
		if err != nil {
			logger.Log().Error("Failed to save transaction", slog.String("sender_id", senderID), slog.String("receiver_id", receiverID))
			return err
		}

		return nil
	})
}
