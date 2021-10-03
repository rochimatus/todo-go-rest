package model

import "time"

type ToDo struct {
	ID        int
	User      User
	UserID    int    `gorm:"type:not null"`
	Title     string `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
