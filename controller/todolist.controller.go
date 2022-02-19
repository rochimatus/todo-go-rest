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

type ToDoListController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type toDoListController struct {
	toDoListService service.ToDoListService
}

func NewToDoListController(toDoListService service.ToDoListService) ToDoListController {
	return &toDoListController{
		toDoListService: toDoListService,
	}
}

func (controller *toDoListController) Create(c *gin.Context) {
	var req request.ToDoListRequest

	err := c.ShouldBind(&req)
	if exception.Error(c, err) {
		return
	}

	toDoList, err := controller.toDoListService.Create(req)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDoListToResponse(toDoList),
		"message": "ToDoList Created Successfully",
	})
}

func (controller *toDoListController) GetAll(c *gin.Context) {
	toDoLists, err := controller.toDoListService.FindAll()
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDoListsToResponses(toDoLists),
		"message": "Get All Data Successfully",
	})
}

func (controller *toDoListController) Get(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	toDoList, err := controller.toDoListService.FindByID(ID)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDoListToResponse(toDoList),
		"message": "Get One Successfully",
	})
}

func (controller *toDoListController) Edit(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	var req request.ToDoListRequest
	err = c.ShouldBind(&req)
	if exception.Error(c, err) {
		return
	}

	toDoList, err := controller.toDoListService.Update(ID, req)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDoListToResponse(toDoList),
		"message": "Edit ToDoList Successfully",
	})
}

func (controller *toDoListController) Delete(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	toDoList, err := controller.toDoListService.Delete(ID)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.ToDoListToResponse(toDoList),
		"message": "Deleted successfully",
	})
}
