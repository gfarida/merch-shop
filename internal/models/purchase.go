package models

import "time"

type Purchase struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"userId" db:"user_id"`
	MerchID   string    `json:"merchId" db:"merch_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
