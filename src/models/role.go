package models

import (
	"db"
)

// Role 角色描述
type Role struct {
	db.Model
	Name        string `json:"name" xml:"name"`
	Description string `json:"description" xml:"description"`
}
