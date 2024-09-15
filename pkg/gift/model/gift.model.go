package giftModel

import (
	userGiftModel "mit-api/pkg/user-gift/model"

	"gorm.io/gorm"
)

type GiftType int
type Status int

type Gift struct {
	gorm.Model
	Name            *string `json:"name" validate:"required,min=2,max=30"`
	Point           *string `json:"point" validate:"required"`
	Status          Status  `json:"status" gorm:"default:1"`
	Quantity        *int    `json:"quantity" validate:"required"`
	UserGifts       []userGiftModel.UserGift
	SpinWheelPrizes []SpinWheelPrize
}

type SpinWheelPrize struct {
	gorm.Model
	GiftID      uint   `json:"gift_id"`
	SpinWheelID uint   `json:"spin_wheel_id"`
	Status      Status `json:"status" gorm:"default:1"`
	Gift        Gift
}
