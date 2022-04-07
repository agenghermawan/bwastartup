package user

import "time"

type Users struct {
	ID           int
	Name         string
	Occupation   string
	Email        string
	passwordHash string
	Avatar       string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
