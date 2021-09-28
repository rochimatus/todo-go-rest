package controller

import (
	"net/http"
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	authenticatedUser, isAuthenticated := controller.authService.Login(credential.Email, credential.Password)
	if isAuthenticated != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  userToCredentialResponse(authenticatedUser),
		"token": controller.jwtService.GenerateToken(authenticatedUser.Email, authenticatedUser.RoleID, true),
	})
}

func (controller *authController) Register(c *gin.Context) {
	var credential request.RegisterRequest

	err := c.ShouldBind(&credential)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	savedUser, err := controller.authService.Register(credential)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user":  userToCredentialResponse(savedUser),
			"token": controller.jwtService.GenerateToken(savedUser.Email, savedUser.Role.ID, true),
		})
	}
}

func userToCredentialResponse(user model.User) response.CredentialResponse {
	return response.CredentialResponse{
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role.ID,
	}
}
