package controller

import (
	"todo-go-rest/helper"
	"todo-go-rest/service"
)

type Controller struct {
	AuthController     AuthController
	RoleController     RoleController
	UserController     UserController
	StatusController   StatusController
	ToDoController     ToDoController
	ToDoListController ToDoListController
}

func CreateController(service *service.Service, helper helper.AuthHelper) *Controller {
	return &Controller{
		AuthController:     NewAuthController(service.AuthService, service.JWTService),
		RoleController:     NewRoleController(service.RoleService),
		UserController:     NewUserController(service.UserService),
		StatusController:   NewStatusController(service.StatusService),
		ToDoController:     NewToDoController(service.ToDoService, service.UserService, helper),
		ToDoListController: NewToDoListController(service.ToDoListService),
	}
}
