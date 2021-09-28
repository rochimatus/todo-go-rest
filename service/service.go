package service

import "todo-go-rest/repository"

type Service struct {
	RoleService RoleService
	AuthService AuthService
	JWTService  JWTService
}

func CreateService(repo *repository.Repository) *Service {

	roleService := NewRoleService(repo.RoleRepository)
	authService := NewAuthService(repo.UserRepository, roleService)
	jwtService := NewJWTService()

	return &Service{
		RoleService: roleService,
		AuthService: authService,
		JWTService:  jwtService,
	}
}
