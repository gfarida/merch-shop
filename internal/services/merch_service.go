package services

import (
	"merch-shop/internal/repository"
	"merch-shop/pkg/errors"
)

func BuyItem(userID, itemName string) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return errors.ErrUserNotFound
	}

	item, err := repository.GetItemByName(itemName)
	if err != nil {
		return errors.ErrItemNotFound
	}

	if user.Balance < item.Price {
		return errors.ErrInsufficientFunds
	}

	user.Balance -= item.Price

	err = repository.AddToInventory(userID, item.ID, 1) // 1 штука товара
	if err != nil {
		return err
	}

	err = repository.UpdateUserBalance(userID, user.Balance) // Исправлено
	if err != nil {
		return err
	}

	err = repository.SavePurchase(userID, item.ID)
	if err != nil {
		return err
	}

	return nil
}
