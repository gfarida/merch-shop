package models

type InventoryItem struct {
	UserID   string `json:"userId" db:"user_id"`
	MerchID  string `json:"merchId" db:"merch_id"`
	Quantity int64  `json:"quantity" db:"quantity"`
}
