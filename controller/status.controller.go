package controller

import (
	"net/http"
	"strconv"
	"todo-go-rest/exception"
	"todo-go-rest/model/request"
	"todo-go-rest/service"

	"github.com/gin-gonic/gin"
)

type StatusController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type statusController struct {
	statusService service.StatusService
}

func NewStatusController(statusService service.StatusService) StatusController {
	return &statusController{
		statusService: statusService,
	}
}

func (controller *statusController) Create(c *gin.Context) {
	var req request.StatusRequest

	err := c.ShouldBind(&req)
	if exception.Error(c, err) {
		return
	}

	status, err := controller.statusService.Create(req)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    status,
		"message": "Status Created Successfully",
	})
}

func (controller *statusController) GetAll(c *gin.Context) {
	statuses, err := controller.statusService.FindAll()
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    statuses,
		"message": "Get All Data Successfully",
	})
}

func (controller *statusController) Get(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	status, err := controller.statusService.FindByID(ID)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    status,
		"message": "Get One Successfully",
	})
}

func (controller *statusController) Edit(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	var req request.StatusRequest
	err = c.ShouldBind(&req)
	if exception.Error(c, err) {
		return
	}

	status, err := controller.statusService.Update(ID, req)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    status,
		"message": "Edit Status Successfully",
	})
}

func (controller *statusController) Delete(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	status, err := controller.statusService.Delete(ID)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    status,
		"message": "Deleted successfully",
	})
}
