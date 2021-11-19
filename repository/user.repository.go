package repository

import (
	"todo-go-rest/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(ID int) (model.User, error)
	Create(user model.User) (model.User, error)
	Delete(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.User, error) {
	var users []model.User
	r.db.Preload("Role").Find(&users)
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) FindByID(ID int) (model.User, error) {
	var user model.User
	r.db.Preload("Role").Find(&user)
	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) Create(user model.User) (model.User, error) {
	r.db.Preload("Role").Find(&user)
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Delete(user model.User) (model.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}

func (r *repository) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) FindByEmail(email string) (model.User, error) {
	var user model.User
	r.db.Preload("Role").Find(&user)
	err := r.db.Where("email = ?", email).First(&user).Error

	return user, err
}
