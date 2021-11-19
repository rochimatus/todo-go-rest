package model

import (
	"time"

	"gorm.io/gorm"
)

type ToDoList struct {
	gorm.Model
	ID        int
	ToDo      ToDo
	ToDoID    int    `gorm:"type:not null"`
	Task      string `gorm:"type:varchar(75);not null"`
	StatusID  int    `gorm:"type:not null;default:1"`
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}
