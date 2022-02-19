package model

import (
	"time"

	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	ID        int
	UserID    int `gorm:"type:not null"`
	User      User
	Title     string `gorm:"type:varchar(50);not null"`
	ToDoLists []ToDoList
	CreatedAt time.Time
	UpdatedAt time.Time
}
