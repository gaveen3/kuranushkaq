package middleware

import (
	"github.com/jinzhu/gorm"

	iris "gopkg.in/kataras/iris.v6"
)

const (
	//CONTEXT *
	CONTEXT string = "DB"
)

type dbMiddleware struct {
	db *gorm.DB
}

func (m *dbMiddleware) Serve(ctx *iris.Context) {
	ctx.Set(CONTEXT, m.db)
	ctx.Next()
}

//NewDBMiddleware *
func NewDBMiddleware(db *gorm.DB) iris.HandlerFunc {
	dbmw := &dbMiddleware{db: db}
	return dbmw.Serve
}
