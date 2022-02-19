package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	StatusCode int `json:"code"`
	Err        error
	Message    string `json:"message"`
}

func ErrorCustom(c *gin.Context, err error, statusCode int) bool {
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return true
	}
	return false
}

func Error(c *gin.Context, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return true
	}
	return false
}

func ErrorWithMessage(c *gin.Context, ae AppError) bool {
	if ae.Err != nil {
		c.AbortWithStatusJSON(ae.StatusCode, gin.H{
			"status":  false,
			"message": ae.Message,
		})
		return true
	}
	return false
}

func (ae *AppError) Error() string {
	return ae.Message
}
