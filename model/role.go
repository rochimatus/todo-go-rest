package model

type Role struct {
	ID          int
	Name        string `gorm:"type:varchar(10);unique;not null"`
	Description string `gorm:"type:varchar(100);"`
}
