package controller

import (
	"net/http"
	"strconv"
	"todo-go-rest/exception"
	"todo-go-rest/helper"
	"todo-go-rest/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) GetAll(c *gin.Context) {
	users, err := controller.userService.FindAll()
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.UsersToUserResponses(users),
		"message": "Get All Data Successfully",
	})
}

func (controller *userController) Get(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	user, err := controller.userService.FindByID(ID)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.UserToUserResponse(user),
		"message": "Get One Successfully",
	})
}

func (controller *userController) Delete(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	user, err := controller.userService.Delete(ID)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.UserToUserResponse(user),
		"message": "Deleted successfully",
	})
}
