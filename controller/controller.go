package controller

import "todo-go-rest/service"

type Controller struct {
	AuthController     AuthController
	RoleController     RoleController
	StatusController   StatusController
	ToDoController     ToDoController
	ToDoListController ToDoListController
}

func CreateController(service *service.Service) *Controller {
	return &Controller{
		AuthController:     NewAuthController(service.AuthService, service.JWTService),
		RoleController:     NewRoleController(service.RoleService),
		StatusController:   NewStatusController(service.StatusService),
		ToDoController:     NewToDoController(service.ToDoService),
		ToDoListController: NewToDoListController(service.ToDoListService),
	}
}
