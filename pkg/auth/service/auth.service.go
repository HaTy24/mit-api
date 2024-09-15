package authService

import (
	"context"
	"log"
	"mit-api/internal/database"
	"mit-api/internal/helpers"
	"mit-api/internal/types"
	tourModel "mit-api/pkg/tour/model"
	userModel "mit-api/pkg/user/model"
	"net/http"
	"time"
)

func Login(email string, password string) *types.OperationResult {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var foundUser userModel.User

	if err := database.DBInstance.WithContext(ctx).Model(&userModel.User{}).Where("email = ?", email).Take(&foundUser); err.Error != nil {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "user not found",
			Code:     "USER_NOT_FOUND",
			HttpCode: http.StatusNotFound,
		}
	}

	var isValidPassword = helpers.CheckPasswordHash(password, *foundUser.Password)
	if !isValidPassword {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "invalid password",
			Code:     "INVALID_PASSWORD",
			HttpCode: http.StatusBadRequest,
		}
	}

	token, err := helpers.TokenGenerator(foundUser.ID, email)
	if err != nil {
		log.Println(err)
	}

	if err := database.DBInstance.WithContext(ctx).Model(&userModel.User{}).Where("email = ?", email).Update("token", token); err.Error != nil {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "update user token failed",
			Code:     "UPDATE_USER_TOKEN_FAILED",
			HttpCode: http.StatusBadRequest,
		}
	}

	// Hide the password by setting it to nil
	foundUser.Password = nil

	foundUser.Token = &token

	return &types.OperationResult{
		Success:  true,
		Data:     &foundUser,
		Message:  "",
		Code:     "",
		HttpCode: http.StatusOK,
	}
}

func SignUp(user userModel.User) *types.OperationResult {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var count int64
	if database.DBInstance == nil {
		log.Fatalf("Database connection is not initialized")
	}

	database.DBInstance.WithContext(ctx).Model(&userModel.User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "user already exists",
			Code:     "USER_ALREADY_EXITS",
			HttpCode: http.StatusBadRequest,
		}
	}

	database.DBInstance.WithContext(ctx).Model(&userModel.User{}).Where("phone_number = ?", user.Phone_Number).Count(&count)
	defer cancel()
	if count > 0 {
		return &types.OperationResult{
			Success:  false,
			Data:     nil,
			Message:  "phone number already in use",
			Code:     "PHONE_NUMBER_ALREADY_IN_USE",
			HttpCode: http.StatusBadRequest,
		}
	}

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	passwordHashed := helpers.HashPassword(*user.Password)
	user.Password = &passwordHashed
	user.Tours = make([]tourModel.Tour, 0)

	if err := database.DBInstance.WithContext(ctx).Model(&userModel.User{}).Create(&user); err.Error != nil {
		log.Println("can not create new user", err)

		return &types.OperationResult{
			Success:  true,
			Data:     nil,
			Message:  err,
			Code:     "",
			HttpCode: http.StatusOK,
		}
	}
	defer cancel()

	return &types.OperationResult{
		Success:  true,
		Data:     nil,
		Message:  "sign up success",
		Code:     "",
		HttpCode: http.StatusOK,
	}
}
