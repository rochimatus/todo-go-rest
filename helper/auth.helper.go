package helper

import (
	"fmt"
	"net/http"
	"todo-go-rest/model"
	"todo-go-rest/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthHelper interface {
	GetCurrentUser(*gin.Context) model.User
}

type authHelper struct {
	userService service.UserService
}

func NewAuthHelper(userService service.UserService) *authHelper {
	return &authHelper{
		userService: userService,
	}
}

func (helper *authHelper) GetCurrentUser(c *gin.Context) (user model.User) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := service.NewJWTService().ValidateToken(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		claimEmail := claims["name"].(string)
		user, err = helper.userService.FindByEmail(claimEmail)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	return user
}
