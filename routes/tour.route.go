package routers

import (
	tour "mit-api/internal/api/tour"

	"github.com/gin-gonic/gin"
)

func TourRoutes(incomingRoutes *gin.RouterGroup) {
	group := incomingRoutes.Group("/tours")
	group.POST("/register", tour.RegisterTour())
	group.GET("", tour.GetTours())
	group.PATCH("", tour.UpdateTour())
	group.DELETE("/:id/cancel", tour.CancelTour())
}
