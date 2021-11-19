package model

type Status struct {
	ID   int
	Name string `gorm:"type:varchar(25);not null"`
}
