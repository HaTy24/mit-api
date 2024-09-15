package rbacModel

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	RoleId       uint `json:"role_id"`
	PermissionId uint `json:"permission_id"`
}
