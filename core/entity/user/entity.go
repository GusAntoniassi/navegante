package user

import "time"

type ID uint64

type User struct {
	ID        ID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
