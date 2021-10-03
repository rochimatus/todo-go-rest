package service

import (
	"errors"
	"strconv"
	"todo-go-rest/model"
	"todo-go-rest/model/request"
	"todo-go-rest/repository"
)

type ToDoService interface {
	FindAll() ([]model.ToDo, error)
	FindByID(ID int) (model.ToDo, error)
	Create(req request.ToDoRequest) (model.ToDo, error)
	Delete(ID int) (model.ToDo, error)
	Update(ID int, req request.ToDoRequest) (model.ToDo, error)
}

type toDoService struct {
	repository repository.ToDoRepository
}

func NewToDoService(repository repository.ToDoRepository) *toDoService {
	return &toDoService{
		repository: repository,
	}
}

func (service *toDoService) FindAll() ([]model.ToDo, error) {
	return service.repository.FindAll()
}

func (service *toDoService) FindByID(ID int) (model.ToDo, error) {
	return service.repository.FindByID(ID)
}

func (service *toDoService) Create(req request.ToDoRequest) (model.ToDo, error) {
	toDo := model.ToDo{
		Title: req.Title,
	}

	return service.repository.Create(toDo)
}

func (service *toDoService) Delete(ID int) (model.ToDo, error) {
	toDo, err := service.repository.FindByID(ID)

	if err != nil {
		return toDo, errors.New("ToDo with ID " + strconv.Itoa(ID) + " is not found")
	}

	return service.repository.Delete(toDo)
}

func (service *toDoService) Update(ID int, req request.ToDoRequest) (model.ToDo, error) {
	toDo, err := service.repository.FindByID(ID)

	if err != nil {
		return toDo, errors.New("ToDo with ID " + strconv.Itoa(ID) + " is not found")
	}

	toDo.Title = req.Title

	return service.repository.Update(toDo)
}
