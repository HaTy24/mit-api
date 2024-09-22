package route

import (
	controller "mit-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	group := incomingRoutes.Group("/auth")
	group.POST("/signup", controller.SignUp())
	group.POST("/login", controller.Login())
}
