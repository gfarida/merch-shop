package models

import "time"

type Merch struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Price     int64      `json:"price" db:"price"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
