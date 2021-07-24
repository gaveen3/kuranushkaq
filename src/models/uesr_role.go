package models

import (
	"db"
)

// UsersRoles 用户和角色关系字段属性描述
type UsersRoles struct {
	db.Model
	UserID      string `json:"user_id" xml:"user_id"`
	RoleID      string `json:"role_id" xml:"role_id"`
	Description string `json:"description" xml:"description"`
}
