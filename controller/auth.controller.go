package controller

import (
	"net/http"
	"todo-go-rest/exception"
	"todo-go-rest/model"
	"todo-go-rest/model/request"
	"todo-go-rest/model/response"
	"todo-go-rest/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (controller *authController) Login(c *gin.Context) {
	var credential request.LoginRequest

	err := c.ShouldBind(&credential)
	if exception.Error(c, err) {
		return
	}

	authenticatedUser, isAuthenticated := controller.authService.Login(credential.Email, credential.Password)
	if isAuthenticated != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"user":   userToCredentialResponse(authenticatedUser),
		"token":  controller.jwtService.GenerateToken(authenticatedUser.Email, authenticatedUser.RoleID, true),
	})
}

func (controller *authController) Register(c *gin.Context) {
	var credential request.RegisterRequest

	err := c.ShouldBind(&credential)
	if exception.Error(c, err) {
		return
	}

	savedUser, err := controller.authService.Register(credential)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"user":   userToCredentialResponse(savedUser),
		"token":  controller.jwtService.GenerateToken(savedUser.Email, savedUser.Role.ID, true),
	})

}

func userToCredentialResponse(user model.User) response.CredentialResponse {
	return response.CredentialResponse{
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role.ID,
	}
}
