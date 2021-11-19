package main

import (
	"todo-go-rest/controller"
	"todo-go-rest/middleware"
	"todo-go-rest/repository"
	"todo-go-rest/service"

	"github.com/gin-gonic/gin"
)

func main() {

	repo := repository.CreateRepository(true)
	// repository.Seeding(repo)
	service := service.CreateService(repo)
	controller := controller.CreateController(service)
	createRoute(controller)
	// var userRepository repository.UserRepository = repository.NewUserRepository(db)
	// var loginService service.AuthService = service.NewAuthService(repo.UserRepository)
	// var jwtService service.JWTService = service.JWTAuthService()
	// var authController controller.AuthController = controller.NewAuthController(loginService, jwtService)
}

func createRoute(controller *controller.Controller) {
	router := gin.Default()

	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.POST("/login", controller.AuthController.Login)
	router.POST("/register", controller.AuthController.Register)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	// authorized := router.Group("/")

	// authorized.Use(middleware.AuthorizeJWT())
	// {
	role := router.Group("/roles")
	role.Use(middleware.Admin())
	{
		role.POST("/", controller.RoleController.Create)
		role.GET("/", controller.RoleController.GetAll)
		role.GET("/:id", controller.RoleController.Get)
		role.PUT("/:id", controller.RoleController.Edit)
		role.DELETE("/:id", controller.RoleController.Delete)
	}
	status := router.Group("/status")
	status.Use(middleware.Admin())
	{
		status.POST("/", controller.StatusController.Create)
		status.GET("/", controller.StatusController.GetAll)
		status.GET("/:id", controller.StatusController.Get)
		status.PUT("/:id", controller.StatusController.Edit)
		status.DELETE("/:id", controller.StatusController.Delete)
	}
	toDo := router.Group("/to-do")
	toDo.Use(middleware.User())
	{
		toDo.POST("/", controller.ToDoController.Create)
		toDo.GET("/", controller.ToDoController.GetAll)
		toDo.GET("/:id", controller.ToDoController.Get)
		toDo.PUT("/:id", controller.ToDoController.Edit)
		toDo.DELETE("/:id", controller.ToDoController.Delete)
	}
	toDoList := router.Group("/to-do-list")
	toDoList.Use(middleware.User())
	{
		toDoList.POST("/", controller.ToDoListController.Create)
		toDoList.GET("/", controller.ToDoListController.GetAll)
		toDoList.GET("/:id", controller.ToDoListController.Get)
		toDoList.PUT("/:id", controller.ToDoListController.Edit)
		toDoList.DELETE("/:id", controller.ToDoListController.Delete)
	}
	router.Run()
}
