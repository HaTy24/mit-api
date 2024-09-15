package routers

import (
	spinwheel "mit-api/internal/api/spin-wheel"

	"github.com/gin-gonic/gin"
)

func SpinWheelRoutes(incomingRoutes *gin.RouterGroup) {
	group := incomingRoutes.Group("/spin-wheels")
	group.POST("/spin", spinwheel.Spin())
}
