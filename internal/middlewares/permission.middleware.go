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

func Permission(permission []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		userID, exist := c.Get("x-user-id")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})

			return
		}

		var user userModel.User
		if err := database.DBInstance.WithContext(ctx).Model(&userModel.User{}).Where("id = ?", userID).Preload("Role.RolePermissions.Permission").First(&user); err.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied 1"})
			c.Abort()
		}

		var userPermission []string
		for i := 0; i < len(user.Role.RolePermissions); i++ {
			userPermission = append(userPermission, *user.Role.RolePermissions[i].Permission.Name)
		}

		for i := 0; i < len(permission); i++ {
			if helpers.Contains(userPermission, permission[i]) {
				c.Next()

				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied 2"})
		c.Abort()
	}
}
