package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/models"
	"github.com/lai0xn/hackiwna-backend/internal/storage"
	uuid "github.com/satori/go.uuid"
)

func IsDoctor() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.FromString(c.MustGet("id").(string))
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Unothorizeed"))
		}
		var user models.User
		db := storage.DB.Where("id = ?", id).Find(&user)
		err = db.Error
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Unothorizeed"))
		}
		fmt.Println(user.IsDoctor)
		if user.IsDoctor == false {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Unothorizeed"))
		}
		c.Next()
	}
}
