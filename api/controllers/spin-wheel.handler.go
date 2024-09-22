package controller

import (
	spinWheelService "mit-api/pkg/spin-wheel/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTags spin
// @Summary spin
// @Description Spin
// @Tags Tours
// @Accept json
// @Produce json
// @Router /spin-wheels/spin [post]
// @Security BearerAuth
func Spin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exist := c.Get("x-user-id")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})

			return
		}
		spinResutl := spinWheelService.SpinProccess(userID.(uint))
		if !spinResutl.Success {
			c.JSON(http.StatusBadRequest, spinResutl)

			return
		}

		c.JSON(http.StatusOK, spinResutl)
	}
}
