package userGiftModel

import (
	"gorm.io/gorm"
)

type Status int

type UserGift struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	GiftID uint   `json:"gift_id"`
	Status Status `json:"status" gorm:"default:1"`
}
