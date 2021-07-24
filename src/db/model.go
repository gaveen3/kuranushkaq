package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

//Model *
type Model struct {
	ID        string     `gorm:"primary_key" json:"id" xml:"id"`
	CreatedAt time.Time  `json:"created_at" xml:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" xml:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at" xml:"deleted_at"`
}

//BeforeCreate ID处理
func (d *Model) BeforeCreate(scope *gorm.Scope) error {
	uuidStr := uuid.NewRandom().String()
	if err := scope.SetColumn("ID", uuidStr); nil != err {
		return err
	}
	return nil
}
