package entity

import "time"

type User struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
