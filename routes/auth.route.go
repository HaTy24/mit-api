package routers

import (
	"mit-api/internal/api/authController"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	group := incomingRoutes.Group("/auth")
	group.POST("/signup", authController.SignUp())
	group.POST("/login", authController.Login())
}
