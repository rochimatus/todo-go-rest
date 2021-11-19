package repository

import (
	"todo-go-rest/model"

	"gorm.io/gorm"
)

type StatusRepository interface {
	FindAll() ([]model.Status, error)
	FindByID(ID int) (model.Status, error)
	Create(status model.Status) (model.Status, error)
	Delete(status model.Status) (model.Status, error)
	Update(status model.Status) (model.Status, error)
}

type statusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *statusRepository {
	return &statusRepository{db}
}

func (r *statusRepository) FindAll() ([]model.Status, error) {
	var statuss []model.Status

	err := r.db.Find(&statuss).Error

	return statuss, err
}

func (r *statusRepository) FindByID(ID int) (model.Status, error) {
	var status model.Status

	err := r.db.First(&status, ID).Error

	return status, err
}

func (r *statusRepository) Create(status model.Status) (model.Status, error) {
	err := r.db.Create(&status).Error

	return status, err
}

func (r *statusRepository) Delete(status model.Status) (model.Status, error) {
	err := r.db.Delete(&status).Error

	return status, err
}

func (r *statusRepository) Update(status model.Status) (model.Status, error) {
	err := r.db.Save(&status).Error

	return status, err
}
