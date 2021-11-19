package model

import (
	"time"

	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	ID        int
	User      User
	UserID    int    `gorm:"type:not null"`
	Title     string `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
