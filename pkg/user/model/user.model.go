package userModel

import (
	tourModel "mit-api/pkg/tour/model"
	userGiftModel "mit-api/pkg/user-gift/model"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_Name   *string `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name    *string `json:"last_name" validate:"required,min=2,max=30"`
	Email        *string `json:"email" validate:"email,required"`
	Phone_Number *string `json:"phone_number" validate:"required"`
	Token        *string `json:"token"`
	Password     *string `json:"password" validate:"required,min=6"`
	Spin_Remaing *int    `json:"spin_remaing"`
	Tours        []tourModel.Tour
	UserGifts    []userGiftModel.UserGift
}
