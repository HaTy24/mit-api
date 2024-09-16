package middlewares

import (
	"context"
	"mit-api/internal/database"
	"mit-api/internal/helpers"
	userModel "mit-api/pkg/user/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		userID, exist := c.Get("x-user-id")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
			c.Abort()
		}

		var user userModel.User
		if err := database.DBInstance.WithContext(ctx).Model(&userModel.User{}).Where("id = ?", userID).Preload("Role").First(&user); err.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied 3"})
			c.Abort()
		}

		for i := 0; i < len(role); i++ {
			if helpers.IsEqual(*user.Role.Name, role[i]) {
				c.Next()

				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied 4"})
		c.Abort()
	}
}
