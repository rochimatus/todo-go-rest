package repository

import (
	"todo-go-rest/model"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll() ([]model.Role, error)
	FindByID(ID int) (model.Role, error)
	Create(role model.Role) (model.Role, error)
	Delete(role model.Role) (model.Role, error)
	Update(role model.Role) (model.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) FindAll() ([]model.Role, error) {
	var roles []model.Role

	err := r.db.Find(&roles).Error

	return roles, err
}

func (r *roleRepository) FindByID(ID int) (model.Role, error) {
	var role model.Role

	err := r.db.First(&role, ID).Error

	return role, err
}

func (r *roleRepository) Create(role model.Role) (model.Role, error) {
	err := r.db.Create(&role).Error

	return role, err
}

func (r *roleRepository) Delete(role model.Role) (model.Role, error) {
	err := r.db.Delete(&role).Error

	return role, err
}

func (r *roleRepository) Update(role model.Role) (model.Role, error) {
	err := r.db.Save(&role).Error

	return role, err
}
