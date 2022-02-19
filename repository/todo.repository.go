package repository

import (
	"todo-go-rest/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ToDoRepository interface {
	FindAll() ([]model.ToDo, error)
	FindByID(ID int) (model.ToDo, error)
	Create(toDo model.ToDo) (model.ToDo, error)
	Delete(toDo model.ToDo) (model.ToDo, error)
	Update(toDo model.ToDo) (model.ToDo, error)
}

type toDoRepository struct {
	db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) *toDoRepository {
	return &toDoRepository{db}
}

func (r *toDoRepository) FindAll() ([]model.ToDo, error) {
	var toDos []model.ToDo

	err := r.db.Find(&toDos).Error
	r.db.Preload(clause.Associations).Find(&toDos)
	r.db.Preload("User.Role").Find(&toDos)

	return toDos, err
}

func (r *toDoRepository) FindByID(ID int) (model.ToDo, error) {
	var toDo model.ToDo

	err := r.db.First(&toDo, ID).Error
	r.db.Preload(clause.Associations).Find(&toDo)
	r.db.Preload("User.Role").Find(&toDo)
	// r.db.Preload("ToDoLists").Find(&toDo)

	return toDo, err
}

func (r *toDoRepository) Create(toDo model.ToDo) (model.ToDo, error) {
	err := r.db.Create(&toDo).Error

	return toDo, err
}

func (r *toDoRepository) Delete(toDo model.ToDo) (model.ToDo, error) {
	err := r.db.Delete(&toDo).Error

	return toDo, err
}

func (r *toDoRepository) Update(toDo model.ToDo) (model.ToDo, error) {
	err := r.db.Save(&toDo).Error

	return toDo, err
}
