package route

import (
	controller "mit-api/api/controllers"
	middleware "mit-api/api/middlewares"

	"github.com/gin-gonic/gin"
)

func TourRoutes(incomingRoutes *gin.RouterGroup) {
	group := incomingRoutes.Group("/tours")
	group.POST("/register", controller.RegisterTour())
	group.GET("", middleware.RoleMiddleware([]string{"Admin", "User"}), middleware.Permission([]string{"READ"}), controller.GetTours())
	group.PATCH("", controller.UpdateTour())
	group.DELETE("/:id/cancel", controller.CancelTour())
}
