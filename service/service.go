package service

import "todo-go-rest/repository"

type Service struct {
	RoleService     RoleService
	AuthService     AuthService
	JWTService      JWTService
	UserService     UserService
	StatusService   StatusService
	ToDoService     ToDoService
	ToDoListService ToDoListService
}

func CreateService(repo *repository.Repository) *Service {

	roleService := NewRoleService(repo.RoleRepository)
	authService := NewAuthService(repo.UserRepository, roleService)
	jwtService := NewJWTService()
	userService := NewUserService(repo.UserRepository, roleService)
	statusService := NewStatusService(repo.StatusRepository)
	toDoService := NewToDoService(repo.ToDoRepository)
	toDoListService := NewToDoListService(repo.ToDoListRepository, toDoService, statusService)

	return &Service{
		RoleService:     roleService,
		AuthService:     authService,
		JWTService:      jwtService,
		UserService:     userService,
		StatusService:   statusService,
		ToDoService:     toDoService,
		ToDoListService: toDoListService,
	}
}

func MadeByUser(id int, userId int) bool {
	return id == userId
}
