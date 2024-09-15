package userService

import (
	"context"
	"mit-api/internal/baseService"
	"mit-api/internal/database"
	"mit-api/internal/types"
	userModel "mit-api/pkg/user/model"
	"net/http"
	"time"
)

type UserService struct {
	*baseService.BaseService[userModel.User]
}

func Initial() *UserService {
	userService := UserService{BaseService: &baseService.BaseService[userModel.User]{DB: database.DBInstance.Model(&userModel.User{})}}

	return &userService
}

func UpdateSpinRemaing(userId uint) *types.OperationResult {
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user userModel.User
	err := Initial().FindById(userId, &user)
	if err != nil {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "user not found",
			Code:     "USER_NOT_FOUND",
			HttpCode: http.StatusBadRequest,
		}
	}

	*user.Spin_Remaing = *user.Spin_Remaing - 1

	err = Initial().Update(&user)
	if err != nil {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "user not found",
			Code:     "USER_NOT_FOUND",
			HttpCode: http.StatusBadRequest,
		}
	}

	return &types.OperationResult{
		Success:  true,
		Data:     nil,
		Message:  "update spin remaing success",
		Code:     "UPDATE_SPIN_REMAING_SUCCESS",
		HttpCode: http.StatusOK,
	}
}
