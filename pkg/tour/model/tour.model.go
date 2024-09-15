package tourModel

import (
	"time"

	"gorm.io/gorm"
)

type Status int

const (
	Waiting Status = iota + 1 // EnumIndex = 1
	Approve                   // EnumIndex = 2
	Reject                    // EnumIndex = 3
	Cancel                    // EnumIndex = 4
)

// String - Creating common behavior - give the type a String function
func (s Status) String() string {
	return [...]string{"Waiting", "Approve", "Reject", "Cancel"}[s-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (s Status) EnumIndex() int {
	return int(s)
}

type Tour struct {
	gorm.Model
	UserId    uint      `json:"user_id"`
	TourdDate time.Time `json:"tourd_date" gorm:"type:date"`
	Status    Status    `json:"status" gorm:"default:1"`
}

type TourWithStatus struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	TourdDate time.Time `json:"tourd_date"`
	Status    string    `json:"status"`
}
