package route

import (
	controller "mit-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func SpinWheelRoutes(incomingRoutes *gin.RouterGroup) {
	group := incomingRoutes.Group("/spin-wheels")
	group.POST("/spin", controller.Spin())
}
