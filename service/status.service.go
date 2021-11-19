package service

import (
	"errors"
	"strconv"
	"todo-go-rest/model"
	"todo-go-rest/model/request"
	"todo-go-rest/repository"
)

type StatusService interface {
	FindAll() ([]model.Status, error)
	FindByID(ID int) (model.Status, error)
	Create(req request.StatusRequest) (model.Status, error)
	Delete(ID int) (model.Status, error)
	Update(ID int, req request.StatusRequest) (model.Status, error)
}

type statusService struct {
	repository repository.StatusRepository
}

func NewStatusService(repository repository.StatusRepository) *statusService {
	return &statusService{
		repository: repository,
	}
}

func (service *statusService) FindAll() ([]model.Status, error) {
	return service.repository.FindAll()
}

func (service *statusService) FindByID(ID int) (model.Status, error) {
	return service.repository.FindByID(ID)
}

func (service *statusService) Create(req request.StatusRequest) (model.Status, error) {
	status := model.Status{
		Name: req.Name,
	}

	return service.repository.Create(status)
}

func (service *statusService) Delete(ID int) (model.Status, error) {
	status, err := service.repository.FindByID(ID)

	if err != nil {
		return status, errors.New("Status with ID " + strconv.Itoa(ID) + " is not found")
	}

	return service.repository.Delete(status)
}

func (service *statusService) Update(ID int, req request.StatusRequest) (model.Status, error) {
	status, err := service.repository.FindByID(ID)

	if err != nil {
		return status, errors.New("Status with ID " + strconv.Itoa(ID) + " is not found")
	}

	status.Name = req.Name

	return service.repository.Update(status)
}
