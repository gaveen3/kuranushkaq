package models

import (
	"db"
)

//RolesResources 角色和资源关系字段属性描述
type RolesResources struct {
	db.Model
	RoleID     string `json:"role_id" xml:"role_id"`
	ResourceID string `json:"resource_id" xml:"resource_id"`
}
