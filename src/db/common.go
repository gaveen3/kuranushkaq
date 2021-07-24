package db

import (
	"middleware"

	"github.com/jinzhu/gorm"
	iris "gopkg.in/kataras/iris.v6"
)

//Config *
type Config struct {
	// Addr *
	Addr string `json:"addr" xml:"addr"`

	// MaxIdleConns *
	MaxIdleConns int `json:"max_idle_conns" xml:"max_idle_conns"`

	// MaxOpenConns *
	MaxOpenConns int `json:"max_open_conns" xml:"max_open_conns"`

	// Username *
	UserName string `json:"username" xml:"username"`

	// Password *
	Password string `json:"password" xml:"password"`

	// DataBase *
	DataBase string `json:"database" xml:"database"`

	//LogMode *
	LogMode bool `json:"log_mode" xml:"log_mode"`
}

//DefaultConfig *
var DefaultConfig = Config{
	Addr: "127.0.0.1:3306",

	// MaxIdleConns *
	MaxIdleConns: 10,
	// MaxOpenConns *
	MaxOpenConns: 100,

	// Username *
	UserName: "root",

	// Password *
	Password: "123456",

	// DataBase *
	DataBase: "test",

	//LogMode *
	LogMode: false,
}

//CtxDB *
func CtxDB(ctx *iris.Context) *gorm.DB {
	return ctx.Get(middleware.CONTEXT).(*gorm.DB)
}

//AutoMigrate *
func AutoMigrate(ctx *iris.Context, schema interface{}) {
	CtxDB(ctx).AutoMigrate(schema)
}
