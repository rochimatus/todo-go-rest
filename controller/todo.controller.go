package controller

import (
	"net/http"
	"strconv"
	"todo-go-rest/exception"
	"todo-go-rest/helper"
	"todo-go-rest/model/request"
	"todo-go-rest/service"

	"github.com/gin-gonic/gin"
)

type ToDoController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type toDoController struct {
	toDoService service.ToDoService
	userService service.UserService
	authHelper  helper.AuthHelper
}

func NewToDoController(toDoService service.ToDoService, userService service.UserService, authHelper helper.AuthHelper) ToDoController {
	return &toDoController{
		toDoService: toDoService,
		userService: userService,
		authHelper:  authHelper,
	}
}

func (controller *toDoController) Create(c *gin.Context) {
	var req request.ToDoRequest

	err := c.ShouldBind(&req)
	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	currentUser := controller.authHelper.GetCurrentUser(c)

	toDo, err := controller.toDoService.Create(req, currentUser)
	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDoToResponse(toDo),
		"message": "ToDo Created Successfully",
	})
}

func (controller *toDoController) GetAll(c *gin.Context) {
	toDos, err := controller.toDoService.FindAll()
	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDosToResponses(toDos),
		"message": "Get All Data Successfully",
	})
}

func (controller *toDoController) Get(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	toDo, err := controller.toDoService.FindByID(ID)
	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDoToResponse(toDo),
		"message": "Get One Successfully",
	})
}

func (controller *toDoController) Edit(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	var req request.ToDoRequest
	err = c.ShouldBind(&req)

	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	currentUser := controller.authHelper.GetCurrentUser(c)
	toDo, err := controller.toDoService.Update(ID, req, currentUser)

	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDoToResponse(toDo),
		"message": "Edit ToDo Successfully",
	})
}

func (controller *toDoController) Delete(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	currentUser := controller.authHelper.GetCurrentUser(c)
	toDo, err := controller.toDoService.Delete(ID, currentUser)

	if exception.ErrorCustom(c, err, http.StatusBadRequest) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDoToResponse(toDo),
		"message": "Deleted successfully",
	})
}
