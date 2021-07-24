package models

import (
	"db"

	iris "gopkg.in/kataras/iris.v6"
)

// User 用户信息字段属性描述
type User struct {
	db.Model
	Name     string `json:"name" xml:"name" gorm:"not null"`
	Email    string `json:"email" xml:"email" gorm:"not null"`
	Password string `json:"password" xml:"password" gorm:"not null"`
}

//SchemaUserInit *
func SchemaUserInit(ctx *iris.Context) {
	db.AutoMigrate(ctx, User{})
}

// InsertUser *
// func InsertUser(ctx *iris.Context, email, password string) (user User) {

// 	db.CtxDB(ctx).Table("sys_users").Where("email = ? and password = ?", email, password).Find(&user)
// 	return
// }

// FindUserByEmailAndPassword *
func FindUserByEmailAndPassword(ctx *iris.Context, email, password string) (user User) {
	db.CtxDB(ctx).Table("sys_users").Where("email = ? and password = ?", email, password).Find(&user)
	return
}
