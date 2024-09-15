package rbacModel

import "gorm.io/gorm"

type Status int

const (
	Active   Status = iota + 1 // EnumIndex = 1
	Inactive                   // EnumIndex = 2
)

// String - Creating common behavior - give the type a String function
func (s Status) String() string {
	return [...]string{"Active", "Inactive"}[s-1]
}

type Role struct {
	gorm.Model
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Status          Status  `json:"status" gorm:"default:1"`
	RolePermissions []RolePermission
}
