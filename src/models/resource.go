package models

import (
	"db"

	iris "gopkg.in/kataras/iris.v6"
)

//Resource 系统资源信息字段属性描述（含菜单和页面元素）
type Resource struct {
	db.Model
	Pid         string     `json:"pid" xml:"pid"`
	SortNum     int        `json:"sort_num" xml:"sort_num"`
	Title       string     `json:"title" xml:"title"`
	I18nID      string     `json:"i18n_id" xml:"i18n_id"`
	TargetURL   string     `json:"target_url" xml:"target_url"`
	Description string     `json:"description" xml:"description"`
	Type        string     `json:"type" xml:"type"`
	IcoClass    string     `json:"ico_class" xml:"ico_class"`
	Children    []Resource `json:"children" xml:"children"`
}

// ViewUserRoleResource 用户、角色、资源关系字段属性描述
type ViewUserRoleResource struct {
	db.Model

	UserID    string `json:"user_id" xml:"user_id"`
	UserName  string `json:"user_name" xml:"user_name"`
	UserEmail string `json:"user_email" xml:"user_email"`

	ResourceID          string `json:"resource_id" xml:"resource_id"`
	ResourcePid         string `json:"resource_pid" xml:"resource_pid"`
	ResourceSortNum     int    `json:"resource_sort_num" xml:"resource_sort_num"`
	ResourceTitle       string `json:"resource_title" xml:"resource_title"`
	ResourceI18nID      string `json:"resource_i18n_id" xml:"resource_i18n_id"`
	ResourceTargetURL   string `json:"resource_target_url" xml:"resource_target_url"`
	ResourceDescription string `json:"resource_description" xml:"resource_description"`
	ResourceType        string `json:"resource_type" xml:"resource_type"`
	ResourceIcoClass    string `json:"resource_ico_class" xml:"resource_ico_class"`

	Children []ViewUserRoleResource `json:"children" xml:"children"`
}

//FindUserMenuByEmail 根据用户名查找对应的菜单
func FindUserMenuByEmail(ctx *iris.Context, email string) []ViewUserRoleResource {
	var usersRolesResources []ViewUserRoleResource
	db.CtxDB(ctx).Table("view_users_roles_resources").Where("resource_type = ? and user_email = ?", "m", email).Group("resource_id").Order("resource_sort_num").Find(&usersRolesResources)
	return usersRolesResources
}
