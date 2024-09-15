package tour

import (
	"log"
	tourModel "mit-api/pkg/tour/model"
	tourService "mit-api/pkg/tour/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTags registerTour
// @Summary registerTour
// @Description Register new Tour
// @Tags Tours
// @Accept json
// @Produce json
// @Param user body request.RegisterTourRequest true "tour"
// @Router /tours/register [post]
// @Security BearerAuth
func RegisterTour() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exist := c.Get("x-user-id")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})

			return
		}

		var tour tourModel.Tour
		if err := c.BindJSON(&tour); err != nil {
			log.Println("err:", err)
		}

		registerTourResult := tourService.RegisterTour(userID.(uint), tour)
		if !registerTourResult.Success {
			c.JSON(http.StatusBadRequest, registerTourResult)

			return
		}

		c.JSON(http.StatusOK, registerTourResult)
	}
}

// CreateTags getTours
// @Summary getTours
// @Description get list of Tours
// @Tags Tours
// @Accept json
// @Produce json
// @Router /tours [get]
// @Security BearerAuth
func GetTours() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exist := c.Get("x-user-id")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})

			return
		}

		getTourResult := tourService.GetTours(userID.(uint))
		if !getTourResult.Success {
			c.JSON(http.StatusBadRequest, getTourResult)

			return
		}

		c.JSON(http.StatusOK, getTourResult)
	}
}

// CreateTags updateTour
// @Summary updateTour
// @Description update Tour
// @Tags Tours
// @Accept json
// @Produce json
// @Param user body request.RegisterTourRequest true "tour"
// @Router /tours [patch]
// @Security BearerAuth
func UpdateTour() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exist := c.Get("x-user-id")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})

			return
		}

		var tour tourModel.Tour
		if err := c.BindJSON(&tour); err != nil {
			log.Println("err:", err)
		}

		updateTourResult := tourService.UpdateTour(userID.(uint), tour)
		if !updateTourResult.Success {
			c.JSON(http.StatusBadRequest, updateTourResult)

			return
		}

		c.JSON(http.StatusOK, updateTourResult)
	}
}

// CreateTags cancelTour
// @Summary cancelTour
// @Description Cancel a specific Tour
// @Tags Tours
// @Accept json
// @Produce json
// @Param id path int true "Tour ID"
// @Router /tours/{id}/cancel [delete]
// @Security BearerAuth
func CancelTour() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exist := c.Get("x-user-id")
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})

			return
		}

		tourId := c.Param("id")

		cancelTourResult := tourService.CancelTour(userID.(uint), tourId)
		if !cancelTourResult.Success {
			c.JSON(http.StatusBadRequest, cancelTourResult)

			return
		}

		c.JSON(http.StatusOK, cancelTourResult)
	}
}
