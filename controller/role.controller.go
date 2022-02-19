package controller

import (
	"net/http"
	"strconv"
	"todo-go-rest/exception"
	"todo-go-rest/helper"
	"todo-go-rest/model"
	"todo-go-rest/model/request"
	"todo-go-rest/model/response"
	"todo-go-rest/service"

	"github.com/gin-gonic/gin"
)

type RoleController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type roleController struct {
	roleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &roleController{
		roleService: roleService,
	}
}

func (controller *roleController) Create(c *gin.Context) {
	var req request.RoleRequest

	err := c.ShouldBind(&req)
	if exception.Error(c, err) {
		return
	}

	role, err := controller.roleService.Create(req)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.RoleToRoleResponse(role),
		"message": "Role Created Successfully",
	})
}

func (controller *roleController) GetAll(c *gin.Context) {
	roles, err := controller.roleService.FindAll()
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    rolesToResponses(roles),
		"message": "Get All Data Successfully",
	})
}

func (controller *roleController) Get(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	role, err := controller.roleService.FindByID(ID)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.RoleToRoleResponse(role),
		"message": "Get One Successfully",
	})
}

func (controller *roleController) Edit(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	var req request.RoleRequest
	err = c.ShouldBind(&req)
	if exception.Error(c, err) {
		return
	}

	role, err := controller.roleService.Update(ID, req)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.RoleToRoleResponse(role),
		"message": "Edit Role Successfully",
	})
}

func (controller *roleController) Delete(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)
	if exception.Error(c, err) {
		return
	}

	role, err := controller.roleService.Delete(ID)
	if exception.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    helper.RoleToRoleResponse(role),
		"message": "Deleted successfully",
	})
}

func roleToResponse(role model.Role) response.RoleResponse {
	return response.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
	}
}

func rolesToResponses(roles []model.Role) (responses []response.RoleResponse) {
	for _, role := range roles {
		responses = append(responses, helper.RoleToRoleResponse(role))
	}
	return responses
}
