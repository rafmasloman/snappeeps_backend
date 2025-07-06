package domain

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Firstname string    `gorm:"column:firstname;not null"`
	Lastname  string    `gorm:"column:lastname;not null"`
	Email     string    `gorm:"column:email;unique;not null"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
