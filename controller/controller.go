package controller

import "todo-go-rest/service"

type Controller struct {
	AuthController AuthController
}

func CreateController(service *service.Service) *Controller {
	return &Controller{
		AuthController: NewAuthController(service.AuthService, service.JWTService),
	}
}
