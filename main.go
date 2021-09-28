package main

import (
	"todo-go-rest/controller"
	"todo-go-rest/middleware"
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
	authorized := router.Group("/")

	authorized.Use(middleware.AuthorizeJWT())
	{
		role := authorized.Group("/roles")
		role.POST("/", controller.RoleController.Create)
		role.GET("/", controller.RoleController.GetAll)
		role.GET("/{id}", controller.RoleController.Get)
		role.PUT("/{id}", controller.RoleController.Edit)
		role.DELETE("/{id}", controller.RoleController.Delete)

		// nested group
		// testing := authorized.Group("testing")
		// testing.GET("/analytics", analyticsEndpoint)
	}

	router.Run()
}
