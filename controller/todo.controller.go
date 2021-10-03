package controller

import (
	"net/http"
	"strconv"
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
}

func NewToDoController(toDoService service.ToDoService) ToDoController {
	return &toDoController{
		toDoService: toDoService,
	}
}

func (controller *toDoController) Create(c *gin.Context) {
	var req request.ToDoRequest

	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	toDo, err := controller.toDoService.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDoToResponse(toDo),
		"status": "ToDo Created Successfully",
	})
}

func (controller *toDoController) GetAll(c *gin.Context) {
	toDos, err := controller.toDoService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDosToResponses(toDos),
		"status": "Get All Data Successfully",
	})
}

func (controller *toDoController) Get(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	toDo, err := controller.toDoService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDoToResponse(toDo),
		"status": "Get One Successfully",
	})
}

func (controller *toDoController) Edit(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var req request.ToDoRequest
	err = c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	toDo, err := controller.toDoService.Update(ID, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDoToResponse(toDo),
		"status": "Edit ToDo Successfully",
	})
}

func (controller *toDoController) Delete(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	toDo, err := controller.toDoService.Delete(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   helper.ToDoToResponse(toDo),
		"status": "Deleted successfully",
	})
}
