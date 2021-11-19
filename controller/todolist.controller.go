package controller

import (
	"net/http"
	"strconv"
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	toDoList, err := controller.toDoListService.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDoListToResponse(toDoList),
		"status": "ToDoList Created Successfully",
	})
}

func (controller *toDoListController) GetAll(c *gin.Context) {
	toDoLists, err := controller.toDoListService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDoListsToResponses(toDoLists),
		"status": "Get All Data Successfully",
	})
}

func (controller *toDoListController) Get(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	toDoList, err := controller.toDoListService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDoListToResponse(toDoList),
		"status": "Get One Successfully",
	})
}

func (controller *toDoListController) Edit(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var req request.ToDoListRequest
	err = c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	toDoList, err := controller.toDoListService.Update(ID, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDoListToResponse(toDoList),
		"status": "Edit ToDoList Successfully",
	})
}

func (controller *toDoListController) Delete(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	toDoList, err := controller.toDoListService.Delete(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDoListToResponse(toDoList),
		"status": "Deleted successfully",
	})
}
