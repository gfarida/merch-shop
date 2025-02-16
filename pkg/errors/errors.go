package errors

import "errors"

var (
	ErrInvalidInput      = errors.New("неверный формат ввода")
	ErrUserNotFound      = errors.New("пользователь не найден")
	ErrInsufficientFunds = errors.New("недостаточно средств на счёте")
	ErrItemNotFound      = errors.New("товар не найден")
	ErrInvalidAuthToken  = errors.New("неверный или просроченный токен")
	ErrInvalidRequest    = errors.New("неправильный запрос")
)
