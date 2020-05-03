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
	//如果为true，则BlockGlobalUpdate会在没有where子句的情况下在更新/删除时生成错误.
	//这是为了防止由于空对象更新/删除而导致最终错误
	db.BlockGlobalUpdate(true)

	// just for debug
	db.LogMode(true)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Token{})
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Comment{})

	return db
}
