package controller

import (
	"net/http"
	"strconv"
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

	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	role, err := controller.roleService.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   roleToResponse(role),
		"status": "Role Created Successfully",
	})
}

func (controller *roleController) GetAll(c *gin.Context) {
	roles, err := controller.roleService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   rolesToResponses(roles),
		"status": "Get All Data Successfully",
	})
}

func (controller *roleController) Get(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	role, err := controller.roleService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   roleToResponse(role),
		"status": "Get One Successfully",
	})
}

func (controller *roleController) Edit(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var req request.RoleRequest
	err = c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	role, err := controller.roleService.Update(ID, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   roleToResponse(role),
		"status": "Edit Role Successfully",
	})
}

func (controller *roleController) Delete(c *gin.Context) {
	str_ID := c.Param("id")
	ID, err := strconv.Atoi(str_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	role, err := controller.roleService.Delete(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   roleToResponse(role),
		"status": "Deleted successfully",
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
		responses = append(responses, roleToResponse(role))
	}
	return responses
}
