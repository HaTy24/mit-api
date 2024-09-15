package surveyModel

import (
	"time"

	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Checkin     time.Time `json:"check-in"`
	Checkout    time.Time `json:"check-out"`
}
