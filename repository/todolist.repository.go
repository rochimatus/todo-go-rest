package repository

import (
	"todo-go-rest/model"

	"gorm.io/gorm"
)

type ToDoListRepository interface {
	FindAll() ([]model.ToDoList, error)
	FindByID(ID int) (model.ToDoList, error)
	Create(toDoList model.ToDoList) (model.ToDoList, error)
	Delete(toDoList model.ToDoList) (model.ToDoList, error)
	Update(toDoList model.ToDoList) (model.ToDoList, error)
}

type toDoListRepository struct {
	db *gorm.DB
}

func NewToDoListRepository(db *gorm.DB) *toDoListRepository {
	return &toDoListRepository{db}
}

func (r *toDoListRepository) FindAll() ([]model.ToDoList, error) {
	var toDoLists []model.ToDoList

	err := r.db.Find(&toDoLists).Error
	r.db.Preload("ToDo").Find(&toDoLists)
	r.db.Preload("Status").Find(&toDoLists)

	return toDoLists, err
}

func (r *toDoListRepository) FindByID(ID int) (model.ToDoList, error) {
	var toDoList model.ToDoList

	err := r.db.First(&toDoList, ID).Error
	r.db.Preload("ToDo").Find(&toDoList)
	r.db.Preload("Status").Find(&toDoList)

	return toDoList, err
}

func (r *toDoListRepository) Create(toDoList model.ToDoList) (model.ToDoList, error) {
	err := r.db.Create(&toDoList).Error

	return toDoList, err
}

func (r *toDoListRepository) Delete(toDoList model.ToDoList) (model.ToDoList, error) {
	err := r.db.Delete(&toDoList).Error

	return toDoList, err
}

func (r *toDoListRepository) Update(toDoList model.ToDoList) (model.ToDoList, error) {
	err := r.db.Save(&toDoList).Error

	return toDoList, err
}
