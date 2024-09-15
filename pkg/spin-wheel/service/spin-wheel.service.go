package spinWheelService

import (
	"context"
	"math/rand"
	"mit-api/internal/baseService"
	"mit-api/internal/database"
	"mit-api/internal/types"
	spinWheelModel "mit-api/pkg/spin-wheel/model"
	userGiftModel "mit-api/pkg/user-gift/model"
	userGiftService "mit-api/pkg/user-gift/service"
	userModel "mit-api/pkg/user/model"
	userService "mit-api/pkg/user/service"
	"net/http"
	"time"
)

type SpinWheelService struct {
	*baseService.BaseService[spinWheelModel.SpinWheel]
}

func Initial() *SpinWheelService {
	spinWheelService := SpinWheelService{BaseService: &baseService.BaseService[spinWheelModel.SpinWheel]{DB: database.DBInstance.Model(&spinWheelModel.SpinWheel{})}}

	return &spinWheelService
}

func SpinProccess(userId uint) *types.OperationResult {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user userModel.User
	err := userService.Initial().FindById(userId, &user)
	if err != nil {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "user not found",
			Code:     "USER_NOT_FOUND",
			HttpCode: http.StatusBadRequest,
		}
	}

	if *user.Spin_Remaing == 0 {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "user not enough spin",
			Code:     "USER_NOT_ENOUGH_SPIN",
			HttpCode: http.StatusBadRequest,
		}
	}

	var foundSpinWheel []spinWheelModel.SpinWheel
	if err := database.DBInstance.WithContext(ctx).Model(&spinWheelModel.SpinWheel{}).Preload("SpinWheelPrizes.Gift").Order("probability desc").Find(&foundSpinWheel); err.Error != nil {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "spin wheel not found",
			Code:     "SPIN_WHEEL_NOT_FOUND",
			HttpCode: http.StatusBadRequest,
		}
	}

	spinResutl := Spin(foundSpinWheel)
	if !spinResutl.Success {
		return spinResutl
	}

	var userGift userGiftModel.UserGift
	if gift, ok := spinResutl.Data.(spinWheelModel.SpinWheel); ok {
		println("GiftID:", gift.SpinWheelPrizes[0].GiftID)
		userGift.GiftID = gift.SpinWheelPrizes[0].GiftID
		userGift.UserID = userId
	} else {
		println("Data không phải là kiểu Gift")
	}

	err = userGiftService.Initial().Create(&userGift)
	if err != nil {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "create user gift error",
			Code:     "CREATE_USER_GIFT_ERROR",
			HttpCode: http.StatusBadRequest,
		}
	}

	if user.Spin_Remaing != nil {
		*user.Spin_Remaing = *user.Spin_Remaing - 1
	}

	err = userService.Initial().Update(&user)
	if err != nil {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "update user error",
			Code:     "UPDATE_USER_ERROR",
			HttpCode: http.StatusBadRequest,
		}
	}

	return spinResutl
}

func Spin(spinWheel []spinWheelModel.SpinWheel) *types.OperationResult {
	// Tính tổng tỉ lệ
	var totalRate float64
	for _, spinWheel := range spinWheel {
		totalRate += *spinWheel.Probability
	}

	// Quay random một số trong khoảng 0 đến tổng tỉ lệ
	spinResult := rand.Float64() * totalRate

	// Duyệt qua danh sách giải thưởng và trả về giải phù hợp với tỉ lệ
	var cumulativeRate float64
	for _, spinWheel := range spinWheel {
		cumulativeRate += *spinWheel.Probability
		if spinResult <= cumulativeRate {
			return &types.OperationResult{
				Success:  true,
				Data:     spinWheel,
				Message:  "spin wheel found",
				Code:     "SPIN_WHEEL_FOUND",
				HttpCode: http.StatusOK,
			}
		}
	}

	return &types.OperationResult{
		Success:  false,
		Data:     nil,
		Message:  "spin wheel error",
		Code:     "SPIN_WHEEL_ERROR",
		HttpCode: http.StatusBadRequest,
	}
}
