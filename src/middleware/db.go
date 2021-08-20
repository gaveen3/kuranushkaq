package middleware

import (
	"github.com/jinzhu/gorm"

	iris "gopkg.in/kataras/iris.v6"
)

const (
	//CONTEXT *
	CONTEXT string = "DB"
)

type dbMw struct {
	db *gorm.DB
}

func (m *dbMw) Serve(ctx *iris.Context) {
	ctx.Set(CONTEXT, m.db)
	ctx.Next()
}

//NewDB *
func NewDB(db *gorm.DB) iris.HandlerFunc {
	dbmw := &dbMw{db: db}
	return dbmw.Serve
}
