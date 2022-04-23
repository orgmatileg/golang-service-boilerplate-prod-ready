package model

import (
	"time"
)

type User struct {
	ID          int
	FullName    string
	PhoneNumber string
	Email       *string
	Status      *int
	PIN         *string
	IsDelete    *bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
