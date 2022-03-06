package model

import (
	"time"

	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	ID         int
	ToDoList   ToDoList
	ToDoListId int    `gorm:"type:not null"`
	Url        string `gorm:"type:varchar(150);not null"`
	Caption    string `gorm:"type:varchar(75);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
