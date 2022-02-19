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
	Create(req request.ToDoRequest, user model.User) (model.ToDo, error)
	Delete(ID int, user model.User) (model.ToDo, error)
	Update(ID int, req request.ToDoRequest, user model.User) (model.ToDo, error)
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

func (service *toDoService) Create(req request.ToDoRequest, user model.User) (model.ToDo, error) {
	toDo := model.ToDo{
		Title:  req.Title,
		UserID: user.ID,
		User:   user,
	}

	return service.repository.Create(toDo)
}

func (service *toDoService) Delete(ID int, user model.User) (model.ToDo, error) {
	toDo, err := service.repository.FindByID(ID)

	if err != nil {
		return model.ToDo{}, errors.New("ToDo with ID " + strconv.Itoa(ID) + " is not found")
	}

	if !madeByUser(toDo.User.ID, user.ID) {
		return model.ToDo{}, errors.New("You can't delete other's toDo")
	}
	return service.repository.Delete(toDo)
}

func (service *toDoService) Update(ID int, req request.ToDoRequest, user model.User) (model.ToDo, error) {
	toDo, err := service.repository.FindByID(ID)

	if err != nil {
		return toDo, errors.New("ToDo with ID " + strconv.Itoa(ID) + " is not found")
	}

	if !madeByUser(toDo.User.ID, user.ID) {
		return model.ToDo{}, errors.New("You can't update other's toDo")
	}
	toDo.Title = req.Title
	return service.repository.Update(toDo)
}

func madeByUser(id int, userId int) bool {
	return id == userId
}
