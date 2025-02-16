package models

type InfoResponse struct {
	Coins       int64              `json:"coins"`
	Inventory   []InventoryItem    `json:"inventory"`
	CoinHistory TransactionHistory `json:"coinHistory"`
}
