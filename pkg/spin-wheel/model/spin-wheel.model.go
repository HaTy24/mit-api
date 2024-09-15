package spinWheelModel

import (
	giftModel "mit-api/pkg/gift/model"

	"gorm.io/gorm"
)

type Status int

type SpinWheel struct {
	gorm.Model
	Name            *string  `json:"name" validate:"required,min=2,max=30"`
	Probability     *float64 `json:"probability" validate:"required"`
	Status          Status   `json:"status" gorm:"default:1"`
	SpinWheelPrizes []giftModel.SpinWheelPrize
}
