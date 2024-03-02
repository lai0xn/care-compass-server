package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lai0xn/hackiwna-backend/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "applicaition/json")
		authHeader := c.GetHeader("Authorization")

		tokenString := strings.Split(authHeader, " ")[1]

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")

			return
		}

		token := utils.ParseToken(tokenString)
		if token == nil {
			fmt.Println(token)
			c.AbortWithStatusJSON(http.StatusBadRequest, "invalid token")
			return
		}
		fmt.Println(token.Valid)
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid  token")
			return

		}
		claims := token.Claims.(jwt.MapClaims)
		c.Set("id", claims["Id"])
		c.Next()
	}
}
