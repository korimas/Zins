package service

import (
	"github.com/jinzhu/gorm"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/common/utils"
	"github.com/zpdev/zins/model"
)

var AuthService = &authService{}

type authService struct {
}

func (service *authService) Login(db *gorm.DB, user *model.User) (*model.User, *errutils.ZinError) {
	var loginUser model.User
	if db.Where("Username = ?", user.Username).First(&loginUser).RecordNotFound() {
		return nil, errutils.UserNotFound(user.Username)
	}
	result, err := utils.VerifyPassphrase(user.Password, []byte(loginUser.Password))
	if err != nil {
		return nil, errutils.PasswordVerifyError()
	}
	if !result {
		return nil, errutils.UserPassError()
	}
	loginUser.Password = ""
	return &loginUser, nil
}

func (service *authService) Logout(db *gorm.DB, user *model.User) *errutils.ZinError {
	return nil
}
