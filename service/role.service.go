package service

import (
	"errors"
	"strconv"
	"todo-go-rest/model"
	"todo-go-rest/model/request"
	"todo-go-rest/repository"
)

type RoleService interface {
	FindAll() ([]model.Role, error)
	FindByID(ID int) (model.Role, error)
	Create(req request.RoleRequest) (model.Role, error)
	Delete(ID int) (model.Role, error)
	Update(ID int, req request.RoleRequest) (model.Role, error)
}

type roleService struct {
	repository repository.RoleRepository
}

func NewRoleService(repository repository.RoleRepository) *roleService {
	return &roleService{
		repository: repository,
	}
}

func (service *roleService) FindAll() ([]model.Role, error) {
	return service.repository.FindAll()
}

func (service *roleService) FindByID(ID int) (model.Role, error) {
	return service.repository.FindByID(ID)
}

func (service *roleService) Create(req request.RoleRequest) (model.Role, error) {
	role := model.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	return service.repository.Create(role)
}

func (service *roleService) Delete(ID int) (model.Role, error) {
	role, err := service.repository.FindByID(ID)

	if err != nil {
		return role, errors.New("Role with ID " + strconv.Itoa(ID) + " is not found")
	}

	return service.repository.Delete(role)
}

func (service *roleService) Update(ID int, req request.RoleRequest) (model.Role, error) {
	role, err := service.repository.FindByID(ID)

	if err != nil {
		return role, errors.New("Role with ID " + strconv.Itoa(ID) + " is not found")
	}

	role.Name = req.Name
	role.Description = req.Description

	return service.repository.Update(role)
}
