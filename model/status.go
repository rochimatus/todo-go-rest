package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	ID        int
	Name      string `gorm:"type:varchar(25);not null"`
	ToDoLists []ToDoList
}
