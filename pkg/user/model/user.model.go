package userModel

import (
	rbacModel "mit-api/pkg/rbac/model"
	tourModel "mit-api/pkg/tour/model"
	userGiftModel "mit-api/pkg/user-gift/model"

	"gorm.io/gorm"
)

type Status int

const (
	Active   Status = iota + 1 // EnumIndex = 1
	Inactive                   // EnumIndex = 2
	Blocked                    // EnumIndex = 3
)

// String - Creating common behavior - give the type a String function
func (s Status) String() string {
	return [...]string{"Active", "Inactive", "Blocked"}[s-1]
}

type User struct {
	gorm.Model
	First_Name   *string `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name    *string `json:"last_name" validate:"required,min=2,max=30"`
	Email        *string `json:"email" validate:"email,required"`
	Phone_Number *string `json:"phone_number" validate:"required"`
	RoleId       uint    `json:"role_id" validate:"required"`
	Status       Status  `json:"status" gorm:"default:1"`
	Token        *string `json:"token"`
	Password     *string `json:"password" validate:"required,min=6"`
	Spin_Remaing *int    `json:"spin_remaing"`
	Tours        []tourModel.Tour
	UserGifts    []userGiftModel.UserGift
	Role         rbacModel.Role
}
