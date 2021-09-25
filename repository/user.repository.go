package repository

import (
	"todo-go-rest/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(ID int) (model.User, error)
	Create(user model.User) (model.User, error)
	Delete(ID int) (model.User, error)
	Update(ID int, user model.User) (model.User, error)
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

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) FindByID(ID int) (model.User, error) {
	var user model.User

	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Delete(ID int) (model.User, error) {
	user, err := r.FindByID(ID)

	err = r.db.Delete(&user).Error

	return user, err
}

func (r *repository) Update(ID int, user model.User) (model.User, error) {
	_, err := r.FindByID(ID)

	err = r.db.Save(&user).Error

	return user, err
}

func (r *repository) FindByEmail(email string) (model.User, error) {
	var user model.User

	err := r.db.Where("email = ?", email).First(&user).Error

	return user, err
}
