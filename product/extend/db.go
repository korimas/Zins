package extend

import (
	"github.com/jinzhu/gorm"
	"github.com/zpdev/zins-extension/orm"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/app"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db != nil {
		return db
	}
	db = orm.New(app.App).DB
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Token{})

	return db
}
