package middleware

import (
	"fmt"
	"net/http"
	"todo-go-rest/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func authorizeJWT(roleId int) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.NewJWTService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			if claims["role"] != roleId {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func Admin() gin.HandlerFunc {
	ADMIN_ID := 1
	return authorizeJWT(ADMIN_ID)
}

func User() gin.HandlerFunc {
	USER_ID := 2
	return authorizeJWT(USER_ID)
}
