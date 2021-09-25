package model

import "time"

type User struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
