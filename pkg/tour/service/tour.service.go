package tourService

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mit-api/internal/cache"
	"mit-api/internal/database"
	"mit-api/internal/types"
	tourModel "mit-api/pkg/tour/model"
	"net/http"
	"strconv"
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

	val := cache.Get(strconv.FormatUint(uint64(userId), 10))

	jsonData, ok := val.([]byte)
	if !ok {
		fmt.Println("Type assertion failed")
	}
	var result map[string]interface{}
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}
	json.Unmarshal(jsonData, &val)
	if val != nil {
		return &types.OperationResult{
			Success:  true,
			Data:     &result,
			Message:  "",
			Code:     "",
			HttpCode: http.StatusOK,
		}
	}

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

	newData, err := json.Marshal(toursWithStatus)
	if err != nil {
		log.Println(err)
	}

	err = cache.Set(strconv.FormatUint(uint64(userId), 10), newData, 5*time.Minute)
	if err != nil {
		log.Println(err)
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
