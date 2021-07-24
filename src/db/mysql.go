package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Gorm 支持
)

//NewMySQL *
func NewMySQL(config Config) (*gorm.DB, error) {

	if len(config.Addr) == 0 {
		config.Addr = DefaultConfig.Addr
	}

	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = DefaultConfig.MaxIdleConns
	}

	if config.MaxOpenConns == 0 {
		config.MaxOpenConns = DefaultConfig.MaxOpenConns
	}

	if len(config.UserName) == 0 {
		config.UserName = DefaultConfig.UserName
	}

	if len(config.Password) == 0 {
		config.Password = DefaultConfig.Password
	}

	if len(config.DataBase) == 0 {
		config.DataBase = DefaultConfig.DataBase
	}

	dbURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&collation=utf8_general_ci&parseTime=True&loc=Asia%%2FShanghai&timeout=30s", config.UserName, config.Password, config.Addr, config.DataBase)

	db, err := gorm.Open("mysql", dbURL)
	if nil != err {
		return nil, err
	}

	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	db.LogMode(config.LogMode)

	if err = db.DB().Ping(); nil != err {
		return nil, err
	}

	return db, nil
}
