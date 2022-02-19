package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          int
	Name        string `gorm:"type:varchar(10);unique;not null"`
	Description string `gorm:"type:varchar(100);"`
}
