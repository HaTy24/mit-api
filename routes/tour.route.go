package routers

import (
	tour "mit-api/internal/api/tourController"
	"mit-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func TourRoutes(incomingRoutes *gin.RouterGroup) {
	group := incomingRoutes.Group("/tours")
	group.POST("/register", tour.RegisterTour())
	group.GET("", middlewares.RoleMiddleware([]string{"Admin", "User"}), middlewares.Permission([]string{"READ"}), tour.GetTours())
	group.PATCH("", tour.UpdateTour())
	group.DELETE("/:id/cancel", tour.CancelTour())
}
