package models

import (
	"db"

	iris "gopkg.in/kataras/iris.v6"
)

//Image 镜像资源信息字段属性描述
type Image struct {
	db.Model
	Name string `json:"name" xml:"name"`
	Type string `json:"type" xml:"tyep"`
	Size int64  `json:"size" xml:"size"`
	Path string `json:"path" xml:"path"`
}

//SchemaImageInit *
func SchemaImageInit(ctx *iris.Context) {
	db.AutoMigrate(ctx, Image{})
}

//FindImagesAll *
func FindImagesAll(ctx *iris.Context) []Image {
	var images []Image
	db.CtxDB(ctx).Find(&images)
	return images
}

//CreateImage *
func CreateImage(ctx *iris.Context, image *Image) {
	db.CtxDB(ctx).Create(image)
}
