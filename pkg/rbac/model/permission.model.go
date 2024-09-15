package rbacModel

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Status          Status  `json:"status" gorm:"default:1"`
	RolePermissions []RolePermission
}
