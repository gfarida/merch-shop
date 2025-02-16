package models

import "time"

type Transaction struct {
	ID         string    `json:"id" db:"id"`
	SenderID   string    `json:"senderId" db:"sender_id"`
	ReceiverID *string   `json:"receiverId,omitempty" db:"receiver_id"`
	TypeID     int       `json:"transactionTypeId" db:"transaction_type_id"`
	Amount     int64     `json:"amount" db:"amount"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

type TransactionType struct {
	ID    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

// История транзакций пользователя
type TransactionHistory struct {
	Received []TransactionReceived `json:"received"`
	Sent     []TransactionSent     `json:"sent"`
}

type TransactionReceived struct {
	FromUser string `json:"fromUser"`
	Amount   int64  `json:"amount"`
}

type TransactionSent struct {
	ToUser string `json:"toUser"`
	Amount int64  `json:"amount"`
}
