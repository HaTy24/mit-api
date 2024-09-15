package request

import "time"

type RegisterTourRequest struct {
	TourdDate time.Time `json:"tourd_date"`
}
