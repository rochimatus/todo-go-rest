package main

import (
	"todo-go-rest/controller"
	"todo-go-rest/repository"
	"todo-go-rest/service"

	"github.com/gin-gonic/gin"
)

func main() {

	repo := repository.CreateRepository()
	service := service.CreateService(repo)
	controller := controller.CreateController(service)
	createRoute(controller)
	// var userRepository repository.UserRepository = repository.NewUserRepository(db)
	// var loginService service.AuthService = service.NewAuthService(repo.UserRepository)
	// var jwtService service.JWTService = service.JWTAuthService()
	// var authController controller.AuthController = controller.NewAuthController(loginService, jwtService)

	// v1 := router.Group("/v1")

	// v1.GET("/", controller.CreateTask)
	// v1.POST("/book", controller.PostBookHandler)
}

func createRoute(controller *controller.Controller) {
	router := gin.Default()

	router.POST("/login", controller.AuthController.Login)
	router.POST("/register", controller.AuthController.Register)
	router.Run()
}
