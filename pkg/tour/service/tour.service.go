package tourService

import (
	"context"
	"log"
	"mit-api/internal/database"
	"mit-api/internal/types"
	tourModel "mit-api/pkg/tour/model"
	"net/http"
	"time"
)

func RegisterTour(userId uint, tour tourModel.Tour) *types.OperationResult {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var tourRepo = database.DBInstance.WithContext(ctx).Model(&tourModel.Tour{})

	var count int64
	tourRepo.Where("user_id = ? AND tourd_date = ?", userId, tour.TourdDate.Format("2006-01-02")).Count(&count)
	if count > 0 {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "user already registered tour",
			Code:     "USER_ALREADY_REGISTERED_SCHEDULE",
			HttpCode: http.StatusBadRequest,
		}
	}

	tour.UserId = userId
	tourRepo.Save(&tour)

	return &types.OperationResult{
		Success:  true,
		Data:     &tour,
		Message:  "",
		Code:     "",
		HttpCode: http.StatusOK,
	}
}

func GetTours(userId uint) *types.OperationResult {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var tourRepo = database.DBInstance.WithContext(ctx).Model(&tourModel.Tour{})

	var tours []tourModel.Tour
	tourRepo.Where("user_id = ?", userId).Find(&tours)
	var toursWithStatus []tourModel.TourWithStatus

	for _, tour := range tours {
		toursWithStatus = append(toursWithStatus, tourModel.TourWithStatus{
			ID:        tour.ID,
			UserId:    tour.UserId,
			TourdDate: tour.TourdDate,
			Status:    tour.Status.String(),
		})
	}

	return &types.OperationResult{
		Success:  true,
		Data:     &toursWithStatus,
		Message:  "",
		Code:     "",
		HttpCode: http.StatusOK,
	}
}

func UpdateTour(userId uint, data tourModel.Tour) *types.OperationResult {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var tourRepo = database.DBInstance.WithContext(ctx).Model(&tourModel.Tour{})

	var tour tourModel.Tour
	err := tourRepo.Where("user_id = ? AND tourd_date = ?", userId, tour.TourdDate.Format("2006-01-02")).Take(&tour)
	if err.Error != nil {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "tour already registered",
			Code:     "SCHEDULE_ALREADY_REGISTERED",
			HttpCode: http.StatusBadRequest,
		}
	}

	database.DBInstance.WithContext(ctx).Model(&tour).Updates(data)

	return &types.OperationResult{
		Success:  true,
		Data:     nil,
		Message:  "update tour success",
		Code:     "",
		HttpCode: http.StatusOK,
	}
}

func CancelTour(userId uint, tourId string) *types.OperationResult {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var tourRepo = database.DBInstance.WithContext(ctx).Model(&tourModel.Tour{})

	var tour tourModel.Tour
	err := tourRepo.Where("user_id = ? AND id = ?", userId, tourId).First(&tour)
	if err.Error != nil {
		log.Println("can not delete", err.Error)

		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "tour not found",
			Code:     "SCHEDULE_NOT_FOUND",
			HttpCode: http.StatusBadRequest,
		}
	}

	if tour.Status != tourModel.Waiting {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "can not cancel tour not waiting",
			Code:     "SCHEDULE_NOT_WAITING",
			HttpCode: http.StatusBadRequest,
		}
	}

	err = database.DBInstance.WithContext(ctx).Model(&tour).Delete(&tour)
	if err.Error != nil {
		log.Println("can not delete", err.Error)

		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "cancel tour failed",
			Code:     "CANCEL_SCHEDULE_FAILED",
			HttpCode: http.StatusBadRequest,
		}
	}

	return &types.OperationResult{
		Success:  true,
		Data:     nil,
		Message:  "cancel tour success",
		Code:     "",
		HttpCode: http.StatusOK,
	}
}
