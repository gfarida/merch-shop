package models

import "time"

type User struct {
	ID           string     `db:"id"`
	Username     string     `db:"username"`
	PasswordHash string     `db:"password_hash"`
	Balance      int64      `db:"balance"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}
