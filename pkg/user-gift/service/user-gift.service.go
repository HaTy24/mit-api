package userGiftService

import (
	"mit-api/internal/baseService"
	"mit-api/internal/database"
	userGiftModel "mit-api/pkg/user-gift/model"
)

type UserService struct {
	*baseService.BaseService[userGiftModel.UserGift]
}

func Initial() *UserService {
	userService := UserService{BaseService: &baseService.BaseService[userGiftModel.UserGift]{DB: database.DBInstance.Model(&userGiftModel.UserGift{})}}

	return &userService
}
