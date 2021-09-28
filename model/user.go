package model

import "time"

type User struct {
	ID        int
	FullName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);unique;not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	Role      Role
	RoleID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
