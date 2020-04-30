package service

import (
	"github.com/jinzhu/gorm"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
)

var UserService = &userService{}

type userService struct {
}

func (service *userService) CreateUser(db *gorm.DB, user *model.User) *errutils.ZinError {
	//if db.First(&user, "username = ?", "zpzhou")
	if db.NewRecord(user) {
		db.Create(user)
	} else {
		return errutils.UserAlreadyExit()
	}
	return nil
}
