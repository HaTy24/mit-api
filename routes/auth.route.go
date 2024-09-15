package routers

import (
	"mit-api/internal/api/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	group := incomingRoutes.Group("/auth")
	group.POST("/signup", auth.SignUp())
	group.POST("/login", auth.Login())
}
